package utils

type BotUtils interface {
	// GenerateEmbedding generates an embedding for the given data.
	GenerateEmbedding(data string) ([]float32, error)

	// AnalyzeImage gets the text from the image and returns the extracted menu items.
	AnalyzeImage(fileBytes []byte, restaurantID uint) (string, error)
}
