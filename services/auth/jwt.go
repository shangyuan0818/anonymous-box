package auth

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"

	"github.com/star-horizon/anonymous-box-saas/database/model"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrTokenNotYet  = errors.New("token not yet valid")
)

// signJwtToken signs a jwt token with the user info.
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
		Issuer:    "auth-service",
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

// parseJwtToken parses a jwt token string and returns the user info.
func (s *AuthServiceImpl) parseJwtToken(ctx context.Context, tokenString string) (*model.User, error) {
	ctx, span := tracer.Start(ctx, "parse-jwt-token")
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"token": tokenString,
	})

	jwtSecret, err := s.SettingRepo.GetByName(ctx, "auth_jwt_secret")
	if err != nil {
		logger.WithError(err).Error("query jwt secret failed")
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		logger.WithError(err).Error("parse token failed")
		return nil, err
	}

	if !token.Valid {
		logger.Error("token invalid")
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.RegisteredClaims)
	if !ok {
		logger.Error("token claims invalid")
		return nil, ErrInvalidToken
	}

	if expire, err := claims.GetExpirationTime(); err != nil {
		logger.WithError(err).Error("get token expire failed")
		return nil, err
	} else if expire.Before(time.Now()) {
		logger.Error("token expired")
		return nil, ErrTokenExpired
	}

	if issuer, err := claims.GetIssuer(); err != nil {
		logger.WithError(err).Error("get token issuer failed")
		return nil, err
	} else if issuer != "auth-service" {
		logger.Error("token issuer invalid")
		return nil, ErrInvalidToken
	}

	if before, err := claims.GetNotBefore(); err != nil {
		logger.WithError(err).Error("get token not before failed")
		return nil, err
	} else if before.After(time.Now()) {
		logger.Error("token not before")
		return nil, ErrTokenNotYet
	}

	userID, err := strconv.ParseUint(claims.ID, 10, 64)
	if err != nil {
		logger.WithError(err).Error("parse user id failed")
		return nil, err
	}

	user, err := s.UserRepo.GetByID(ctx, uint(userID))
	if err != nil {
		logger.WithError(err).Error("query user failed")
		return nil, err
	}

	return user, nil
}
