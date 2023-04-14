package auth

import (
	"bytes"
	"context"
	"errors"
	"github.com/star-horizon/anonymous-box-saas/database/model"
	repo2 "github.com/star-horizon/anonymous-box-saas/database/repo"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.uber.org/fx"
	"golang.org/x/crypto/bcrypt"

	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	verifyapi "github.com/star-horizon/anonymous-box-saas/kitex_gen/api"
	"github.com/star-horizon/anonymous-box-saas/kitex_gen/api/verifyservice"
)

var tracer = otel.Tracer("auth-service")

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct {
	fx.In
	SettingRepo     repo2.SettingRepo
	UserRepo        repo2.UserRepo
	VerifySvcClient verifyservice.Client
}

// NewAuthServiceImpl creates a new AuthServiceImpl.
func NewAuthServiceImpl(impl AuthServiceImpl) api.AuthService {
	return &impl
}

// UsernameAuth implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) UsernameAuth(ctx context.Context, req *api.UsernameAuthRequest) (*api.AuthToken, error) {
	ctx, span := tracer.Start(ctx, "username-auth")
	defer span.End()

	logger := logrus.WithContext(ctx).WithFields(logrus.Fields{
		"user.username": req.GetUsername(),
	})

	user, err := s.UserRepo.GetByUsername(ctx, req.GetUsername())
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
		"user.email": req.GetEmail(),
	})

	user, err := s.UserRepo.GetByEmail(ctx, req.GetEmail())
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

	resp := &api.AuthToken{
		Token: tokenString,
	}

	return resp, nil
}

var (
	ErrInvalidVerificationCode = errors.New("invalid verification code")
	ErrRegisterNotAllowed      = errors.New("register is not allowed")
)

// Register implements the AuthServiceImpl interface.
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
		if resp, err := s.VerifySvcClient.VerifyEmail(ctx, &verifyapi.VerifyEmailRequest{
			Email: req.GetEmail(),
			Code:  req.GetVerificationCode(),
		}); err != nil {
			logger.WithError(err).Error("verify email failed, verify service error")
			return nil, err
		} else if !resp.GetOk() {
			logger.WithError(ErrInvalidVerificationCode).Error("verify email failed, invalid verification status")
			return nil, ErrInvalidVerificationCode
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
	if resp, err := s.VerifySvcClient.VerifyEmail(ctx, &verifyapi.VerifyEmailRequest{
		Email: req.GetEmail(),
		Code:  req.GetVerificationCode(),
	}); err != nil {
		logger.WithError(err).Error("verify email failed")
		return nil, err
	} else if !resp.GetOk() {
		logger.WithError(ErrInvalidVerificationCode).Error("verify email failed")
		return nil, ErrInvalidVerificationCode
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
		Uid:      uint32(user.ID),
		Username: user.Username,
	}, nil
}
