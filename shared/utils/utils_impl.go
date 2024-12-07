package utils

import (
	"errors"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sirupsen/logrus"
)

type UtilsImpl struct{}

// GenerateNanoID implements Utils.
func (u *UtilsImpl) GenerateNanoID() (string, error) {
	const length = 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	id, err := gonanoid.Generate(charset, length)
	if err != nil {
		logrus.WithError(err).Error("*** [UtilsImpl] Error generating nanoid")
		return "", errors.New("error generating nanoid")
	}

	return id, nil
}

// ParseDateTimeToString implements Utils.
func (u *UtilsImpl) ParseDateTimeToString(date time.Time) (*string, error) {
	birthDate := date.Format("2006-01-02")
	if birthDate == "" {
		logrus.Error("*** [UtilsImpl] Error parsing birth date")
		return nil, errors.New("error parsing birth date")
	}

	return &birthDate, nil
}

// ParseStringToDateTime implements Utils.
func (u *UtilsImpl) ParseStringToDateTime(date string) (*time.Time, error) {
	birthDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		logrus.WithError(err).Error("*** [UtilsImpl] Error parsing birth date")
		return nil, err
	}

	return &birthDate, nil
}

func NewUtilsImpl() Utils {
	return &UtilsImpl{}
}
