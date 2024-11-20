package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioUtilsImpl struct {
	twilio *twilio.RestClient
}

// SendWspMessage implements TwilioUtils.
func (t *TwilioUtilsImpl) SendWspMessage(to string, from string, message string) error {

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	resp, err := t.twilio.Api.CreateMessage(params)
	if err != nil {
		logrus.WithError(err).Error("*** [TwilioUtilsImpl] failed to send message")
		return fmt.Errorf("failed to send message: %v", err)
	}

	logrus.WithField("response", resp).Info("message sent")

	return nil
}

func NewTwilioUtilsImpl(twilio *twilio.RestClient) TwilioUtils {
	return &TwilioUtilsImpl{twilio: twilio}
}
