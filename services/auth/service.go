package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"

	"github.com/star-horizon/anonymous-box-saas/internal/database/model"
	"github.com/star-horizon/anonymous-box-saas/internal/database/repo"
	"github.com/star-horizon/anonymous-box-saas/services/auth/kitex_gen/api"
)

var tracer = otel.Tracer("auth-service")

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	fx.In
	SettingRepo repo.SettingRepo
	UserRepo    repo.UserRepo
}

// NewAuthServiceImpl creates a new AuthServiceImpl.
func NewAuthServiceImpl(impl AuthServiceImpl) api.AuthService {
	return &impl
}

func (s *AuthServiceImpl) signJwtToken(user *model.User) (string, error) {
	ctx, span := tracer.Start(context.Background(), "sign-jwt-token")
	defer span.End()

	now := time.Now()

	logger := logrus.WithFields(logrus.Fields{
		"user.id":       user.ID,
		"user.username": user.Username,
		"user.email":    user.Email,
		"issued_at":     now,
		"not_before":    now,
		"expires_at":    now.Add(time.Hour * 24),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "authservice",
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24)),
		Subject:   user.Username,
		ID:        strconv.FormatUint(uint64(user.ID), 10),
	})

	jwtSecret, err := s.SettingRepo.GetByName(ctx, "auth_jwt_secret")
	if err != nil {
		logger.WithError(err).Error("query jwt secret failed")
		return "", err
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return "", err
	}

	return tokenString, nil
}

// UsernameAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) UsernameAuth(ctx context.Context, req *api.UsernameAuthRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "username-auth")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"username": req.GetUsername(),
	})

	user, err := s.UserRepo.GetByUsername(ctx, req.GetUsername())
	if err != nil {
		logger.WithError(err).Error("query user failed")
		return nil, err
	}

	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	resp := &api.AuthToken{
		Token: tokenString,
	}

	return resp, nil
}

// EmailAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) EmailAuth(ctx context.Context, req *api.EmailAuthRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "email-auth")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"email": req.GetEmail(),
	})

	user, err := s.UserRepo.GetByEmail(ctx, req.GetEmail())
	if err != nil {
		logger.WithError(err).Error("query user failed")
		return nil, err
	}

	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	resp := &api.AuthToken{
		Token: tokenString,
	}

	return resp, nil
}

// Register implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) Register(ctx context.Context, req *api.RegisterRequest) (resp *api.AuthToken, err error) {
	// TODO: Your code here...
	return
}

// ChangePassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ChangePassword(ctx context.Context, req *api.ChangePasswordRequest) (resp *api.AuthToken, err error) {
	// TODO: Your code here...
	return
}

// ResetPassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ResetPassword(ctx context.Context, req *api.ResetPasswordRequest) (resp *api.AuthToken, err error) {
	// TODO: Your code here...
	return
}

// GetServerAuthData implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetServerAuthData(ctx context.Context, req *api.AuthToken) (resp *api.ServerAuthDataResponse, err error) {
	// TODO: Your code here...
	return
}
