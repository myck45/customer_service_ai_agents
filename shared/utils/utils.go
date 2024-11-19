package utils

import "time"

type Utils interface {
	ParseStringToDateTime(date string) (*time.Time, error)
	ParseDateTimeToString(date time.Time) (*string, error)
}
