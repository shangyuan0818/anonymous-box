package auth

import (
	"bytes"
	"context"
	"errors"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	verifyapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/verifyservice"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
)

var tracer = otel.Tracer("auth-service")

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	fx.In
	SettingRepo     repo.SettingRepo
	UserRepo        repo.UserRepo
	VerifySvcClient verifyservice.Client
}

// NewAuthServiceImpl creates a new AuthServiceImpl.
func NewAuthServiceImpl(impl AuthServiceImpl) api.AuthService {
	return &impl
}

var (
	ErrNoAuthCredential = errors.New("no auth credential")
)

// Auth implements the api.AuthService interface.
func (s *AuthServiceImpl) Auth(ctx context.Context, req *api.AuthRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "username-auth")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.username": req.GetUsername(),
	})

	var (
		user *model.User
		err  error
	)
	switch true {
	case lo.IsNotEmpty(req.GetUsername()):
		user, err = s.UserRepo.GetByUsername(ctx, req.GetUsername())
	case lo.IsNotEmpty(req.GetEmail()):
		user, err = s.UserRepo.GetByEmail(ctx, req.GetEmail())
	default:
		return nil, ErrNoAuthCredential
	}
	if err != nil {
		logger.WithError(err).Error("query user failed")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		logger.WithError(err).Error("password not match")
		return nil, err
	}

	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &api.AuthToken{
		Token: tokenString,
	}, nil
}

var (
	ErrRegisterNotAllowed = errors.New("register is not allowed")
)

// Register implements the api.AuthService interface.
func (s *AuthServiceImpl) Register(ctx context.Context, req *api.RegisterRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "register")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.username": req.GetUsername(),
		"user.email":    req.GetEmail(),
	})

	if allowRegister, err := s.SettingRepo.GetBoolByName(ctx, "auth_allow_register"); err != nil {
		logger.WithError(err).Error("query setting failed")
		return nil, err
	} else if !allowRegister {
		return nil, ErrRegisterNotAllowed
	}

	// check verification code if require
	if requireEmailVerify, err := s.SettingRepo.GetBoolByName(ctx, "auth_require_email_verify"); err != nil {
		logger.WithError(err).Error("query setting failed")
		return nil, err
	} else if requireEmailVerify {
		if _, err := s.VerifySvcClient.VerifyEmail(ctx, &verifyapi.VerifyEmailRequest{
			Email: req.GetEmail(),
			Code:  req.GetVerificationCode(),
		}); err != nil {
			logger.WithError(err).Error("verify email failed, verify service error")
			return nil, err
		}
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		logger.WithError(err).Error("hash password failed")
		return nil, err
	}

	// create user
	user := &model.User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: string(hashedPassword),
	}
	if err := s.UserRepo.Create(ctx, user); err != nil {
		logger.WithError(err).Error("create user failed")
		return nil, err
	}

	// sign token
	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &api.AuthToken{
		Token: tokenString,
	}, nil
}

var (
	ErrOldPasswordNotMatch = errors.New("old password not match")
	ErrSamePassword        = errors.New("new password is same as old password")
)

// ChangePassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ChangePassword(ctx context.Context, req *api.ChangePasswordRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "change-password")
	defer span.End()

	// parse token
	user, err := s.parseJwtToken(ctx, req.Token)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("parse token failed")
		return nil, err
	}

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.id":       user.ID,
		"user.username": user.Username,
		"user.email":    user.Email,
	})

	// check old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetOldPassword())); err != nil {
		logger.WithError(err).Error("old password not match")
		return nil, ErrOldPasswordNotMatch
	}

	// hash new password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetNewPassword()), bcrypt.DefaultCost)
	if err != nil {
		logger.WithError(err).Error("hash password failed")
		return nil, err
	}

	// check new password is same as old password
	if bytes.Equal(hashedNewPassword, []byte(user.Password)) {
		logger.WithError(err).Error("new password is same as old password")
		return nil, ErrSamePassword
	}

	// update password
	user.Password = string(hashedNewPassword)
	if err := s.UserRepo.Update(ctx, user); err != nil {
		logger.WithError(err).Error("update user failed")
		return nil, err
	}

	// sign new token
	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &api.AuthToken{
		Token: tokenString,
	}, nil
}

// ResetPassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ResetPassword(ctx context.Context, req *api.ResetPasswordRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "reset-password")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.email": req.GetEmail(),
	})

	// check verification code
	if _, err := s.VerifySvcClient.VerifyEmail(ctx, &verifyapi.VerifyEmailRequest{
		Email: req.GetEmail(),
		Code:  req.GetVerificationCode(),
	}); err != nil {
		logger.WithError(err).Error("verify email failed")
		return nil, err
	}

	// get user
	user, err := s.UserRepo.GetByEmail(ctx, req.GetEmail())
	if err != nil {
		logger.WithError(err).Error("get user failed")
		return nil, err
	}

	// hash new password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetNewPassword()), bcrypt.DefaultCost)
	if err != nil {
		logger.WithError(err).Error("hash password failed")
		return nil, err
	}

	// update password
	user.Password = string(hashedNewPassword)
	if err := s.UserRepo.Update(ctx, user); err != nil {
		logger.WithError(err).Error("update user failed")
		return nil, err
	}

	// sign new token
	tokenString, err := s.signJwtToken(user)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &api.AuthToken{
		Token: tokenString,
	}, nil
}

// GetServerAuthData implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetServerAuthData(ctx context.Context, req *api.AuthToken) (*api.ServerAuthDataResponse, error) {
	ctx, span := tracer.Start(ctx, "get-server-auth-data")
	defer span.End()

	// parse token
	user, err := s.parseJwtToken(ctx, req.Token)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("parse token failed")
		return nil, err
	}

	return &api.ServerAuthDataResponse{
		Uid: user.ID,
		CreatedAt: &base.Timestamp{
			Seconds: user.CreatedAt.Unix(),
			Nanos:   int32(user.CreatedAt.Nanosecond()),
		},
		UpdatedAt: &base.Timestamp{
			Seconds: user.UpdatedAt.Unix(),
			Nanos:   int32(user.UpdatedAt.Nanosecond()),
		},
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
