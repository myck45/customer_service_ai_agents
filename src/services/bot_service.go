package services

type BotService interface {
	GenerateEmbedding(data string) ([]float32, error)
	ChatCompletion(data string) (string, error)
	SystemPrompt(aditionalData string) (string, error)
}
