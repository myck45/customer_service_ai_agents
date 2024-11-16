package providers

import (
	"os"

	"github.com/twilio/twilio-go"
)

func NewTwilioClient() *twilio.RestClient {
	acountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: acountSid,
		Password: authToken,
	})

	return client
}
