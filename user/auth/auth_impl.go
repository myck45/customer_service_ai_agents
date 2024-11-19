package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthImpl struct{}

// GenerateToken implements Auth.
func (a *AuthImpl) GenerateToken(id uint, email string, role string) (string, error) {

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		logrus.WithError(err).Error("*** [AuthImpl] Error signing token")
		return "", err
	}

	return signedToken, nil
}

func NewAuth() Auth {
	return &AuthImpl{}
}
