package auth

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type BcryptImpl struct{}

// ComparePassword implements Bcrypt.
func (b *BcryptImpl) ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		logrus.WithError(err).Error("*** [BcryptImpl] Error comparing password")
		return err
	}

	return nil
}

// HashPassword implements Bcrypt.
func (b *BcryptImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.WithError(err).Error("*** [BcryptImpl] Error hashing password")
		return "", err
	}

	return string(hashedPassword), nil
}

func NewBcryptImpl() Bcrypt {
	return &BcryptImpl{}
}
