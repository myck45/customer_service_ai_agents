package utils

type BotUtils interface {
	GenerateEmbedding(data string) ([]float32, error)
}
