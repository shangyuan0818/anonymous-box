package dash_auth

import (
	"bytes"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"github.com/star-horizon/anonymous-box-saas/database/model"
	"github.com/star-horizon/anonymous-box-saas/database/repo"
	"github.com/star-horizon/anonymous-box-saas/internal/jwt"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/base"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	verifyapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/dash"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/dash/verifyservice"
)

var tracer = otel.Tracer("dash-auth-service")
var validate *validator.Validate

const ServiceName = "dash-auth-service"

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	fx.In
	SettingRepo     repo.SettingRepo
	UserRepo        repo.UserRepo
	VerifySvcClient verifyservice.Client
	JwtSvc          jwt.Service
}

// NewAuthServiceImpl creates a new AuthServiceImpl.
func NewAuthServiceImpl(impl AuthServiceImpl) dash.AuthService {
	return &impl
}

var (
	ErrInvalidCredential = errors.New("invalid credential")
)

// Login implements the dash.AuthService interface.
func (s *AuthServiceImpl) Login(ctx context.Context, req *dash.LoginRequest) (*dash.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "username-auth")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.credential": req.GetCredential(),
	})

	var (
		user *model.User
		err  error
	)

	switch nil {
	case validate.Var(req.GetCredential(), "required,email"):
		span.AddEvent("Login by email")
		logger.Debugf("Login by email: %s", req.GetCredential())

		user, err = s.UserRepo.GetByEmail(ctx, req.GetCredential())
	case validate.Var(req.GetCredential(), "required,alphanum"):
		span.AddEvent("Login by username")
		logger.Debugf("Login by username: %s", req.GetCredential())

		user, err = s.UserRepo.GetByUsername(ctx, req.GetCredential())
	default:
		logger.WithError(ErrInvalidCredential).Errorf("invalid credential: %s", req.GetCredential())
		return nil, ErrInvalidCredential
	}
	if err != nil {
		logger.WithError(err).Error("query user failed")
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		logger.WithError(err).Error("password not match")
		return nil, err
	}

	tokenString, err := s.JwtSvc.GenerateToken(ctx, user.ID)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &dash.AuthToken{
		Token: tokenString,
	}, nil
}

var (
	ErrRegisterNotAllowed = errors.New("register is not allowed")
)

// Register implements the dash.AuthService interface.
func (s *AuthServiceImpl) Register(ctx context.Context, req *dash.RegisterRequest) (*dash.AuthToken, error) {
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

	// check verification code if you require
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
	tokenString, err := s.JwtSvc.GenerateToken(ctx, user.ID)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &dash.AuthToken{
		Token: tokenString,
	}, nil
}

var (
	ErrOldPasswordNotMatch = errors.New("old password not match")
	ErrSamePassword        = errors.New("new password is same as old password")
)

// ChangePassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ChangePassword(ctx context.Context, req *dash.ChangePasswordRequest) (*dash.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "change-password")
	defer span.End()

	// parse token
	userId, err := s.JwtSvc.ParseToken(ctx, req.Token)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("parse token failed")
		return nil, err
	}

	user, err := s.UserRepo.GetByID(ctx, userId)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("get user failed")
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
	if err := s.UserRepo.UpdateByID(ctx, user.ID, user); err != nil {
		logger.WithError(err).Error("update user failed")
		return nil, err
	}

	// sign new token
	tokenString, err := s.JwtSvc.GenerateToken(ctx, user.ID)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &dash.AuthToken{
		Token: tokenString,
	}, nil
}

// ResetPassword implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ResetPassword(ctx context.Context, req *dash.ResetPasswordRequest) (*dash.AuthToken, error) {
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
	if err := s.UserRepo.UpdateByID(ctx, user.ID, user); err != nil {
		logger.WithError(err).Error("update user failed")
		return nil, err
	}

	// sign new token
	tokenString, err := s.JwtSvc.GenerateToken(ctx, user.ID)
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return nil, err
	}

	return &dash.AuthToken{
		Token: tokenString,
	}, nil
}

// GetServerAuthData implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GetServerAuthData(ctx context.Context, req *dash.AuthToken) (*dash.ServerAuthDataResponse, error) {
	ctx, span := tracer.Start(ctx, "get-server-auth-data")
	defer span.End()

	// parse token
	userId, err := s.JwtSvc.ParseToken(ctx, req.Token)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("parse token failed")
		return nil, err
	}

	user, err := s.UserRepo.GetByID(ctx, userId)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("get user failed")
		return nil, err
	}

	return &dash.ServerAuthDataResponse{
		Id: user.ID,
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
