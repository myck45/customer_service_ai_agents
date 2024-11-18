package utils

type TwilioUtils interface {
	SendWspMessage(to string, from string, message string) error
}
