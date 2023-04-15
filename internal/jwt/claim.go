package jwt

import "github.com/golang-jwt/jwt/v5"

type userClaims struct {
	jwt.RegisteredClaims
	UserID uint64 `json:"user_id"`
}
