package utils

import "github.com/golang-jwt/jwt/v5"

type AuthUtils interface {
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}
