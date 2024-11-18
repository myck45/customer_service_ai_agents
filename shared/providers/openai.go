package providers

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

func NewOpenAIClient() *openai.Client {
	var openAIApiKey = os.Getenv("OPENAI_API_KEY")

	config := openai.DefaultConfig(openAIApiKey)
	config.BaseURL = "https://models.inference.ai.azure.com"
	client := openai.NewClientWithConfig(config)

	return client
}
