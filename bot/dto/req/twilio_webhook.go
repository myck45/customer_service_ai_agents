package req

type TwilioWebhook struct {
	To   string `json:"To" binding:"required"`
	From string `json:"From" binding:"required"`
	Body string `json:"Body" binding:"required"`
}
