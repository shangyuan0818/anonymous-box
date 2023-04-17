package jwt

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
	ErrTokenNotYet  = errors.New("token not yet valid")
)

const Issuer = "anonymous-box-saas-jwt"

// GenerateToken signs a jwt token with the user info.
func (s *service) GenerateToken(ctx context.Context, userId uint64) (string, error) {
	ctx, span := tracer.Start(ctx, "generate-token", trace.WithAttributes(
		attribute.Int64("user.id", int64(userId)),
	))
	defer span.End()

	now := time.Now()

	logger := logrus.WithFields(logrus.Fields{
		"user.id":    userId,
		"issued_at":  now,
		"not_before": now,
		"expires_at": now.Add(time.Hour * 24),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.expire)),
			ID:        strconv.FormatUint(userId, 10),
			Subject:   fmt.Sprintf("user-%d", userId),
			Audience:  []string{},
		},
	})

	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		logger.WithError(err).Error("sign token failed")
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses a jwt token string and returns the user info.
func (s *service) ParseToken(ctx context.Context, tokenString string) (uint64, error) {
	ctx, span := tracer.Start(ctx, "parse-token", trace.WithAttributes(
		attribute.String("token", tokenString),
	))
	defer span.End()

	logger := logrus.WithFields(logrus.Fields{
		"token": tokenString,
	})

	token, err := jwt.ParseWithClaims(tokenString, &userClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	}, jwt.WithIssuer(Issuer))
	if err != nil {
		logger.WithError(err).Error("parse token failed")
		return 0, err
	}

	if !token.Valid {
		logger.Error("token invalid")
		return 0, ErrInvalidToken
	}

	claims, ok := token.Claims.(*userClaims)
	if !ok {
		logger.Error("token claims invalid")
		return 0, ErrInvalidToken
	}

	if expire, err := claims.GetExpirationTime(); err != nil {
		logger.WithError(err).Error("get token expire failed")
		return 0, err
	} else if expire.Before(time.Now()) {
		logger.Error("token expired")
		return 0, ErrTokenExpired
	}

	if issuer, err := claims.GetIssuer(); err != nil {
		logger.WithError(err).Error("get token issuer failed")
		return 0, err
	} else if issuer != Issuer {
		logger.Error("token issuer invalid")
		return 0, ErrInvalidToken
	}

	if before, err := claims.GetNotBefore(); err != nil {
		logger.WithError(err).Error("get token not before failed")
		return 0, err
	} else if before.After(time.Now()) {
		logger.Error("token not before")
		return 0, ErrTokenNotYet
	}

	return claims.UserID, nil
}
