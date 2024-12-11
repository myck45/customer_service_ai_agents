package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthUtilsImpl struct{}

// ValidateToken validates and parses a JWT token
func (a *AuthUtilsImpl) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Validate input
	if tokenString == "" {
		logrus.Error("Token is empty")
		return nil, fmt.Errorf("token is empty")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		logrus.Error("JWT_SECRET is not set")
		return nil, fmt.Errorf("jwt secret is not configured")
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	},
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithLeeway(5*time.Minute))

	// Check parsing errors
	if err != nil {
		logrus.WithError(err).Error("Error parsing token")
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	// Validate token
	if !token.Valid {
		logrus.Error("Token is not valid")
		return nil, fmt.Errorf("token is not valid")
	}

	// Extract and validate claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logrus.Error("Invalid token claims")
		return nil, fmt.Errorf("invalid token claims")
	}

	// Validate expiration
	if exp, ok := claims["exp"].(float64); ok {
		if time.Now().Unix() > int64(exp) {
			logrus.Error("Token has expired")
			return nil, fmt.Errorf("token has expired")
		}
	} else {
		logrus.Error("Missing expiration claim")
		return nil, fmt.Errorf("missing expiration claim")
	}

	// Additional claim validations
	if id, ok := claims["id"].(float64); !ok || id <= 0 {
		logrus.Error("Invalid user ID claim")
		return nil, fmt.Errorf("invalid user ID claim")
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		logrus.Error("Invalid or missing email claim")
		return nil, fmt.Errorf("invalid or missing email claim")
	}

	role, ok := claims["role"].(string)
	if !ok || role == "" {
		logrus.Error("Invalid or missing role claim")
		return nil, fmt.Errorf("invalid or missing role claim")
	}

	return claims, nil
}

func NewAuthUtilsImpl() AuthUtils {
	return &AuthUtilsImpl{}
}
