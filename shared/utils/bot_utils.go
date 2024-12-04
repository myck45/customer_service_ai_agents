package utils

type BotUtils interface {
	// GenerateEmbedding generates an embedding for the given data.
	GenerateEmbedding(data string) ([]float32, error)

	// BotToolsHandler Wraper for handle funcion calls
	BotToolsHandler(functionName string, data string, restaurantID uint) error

	// HandleGetUserOrder Wraper for handle GetUserOrder function
	HandleGetUserOrder(data string, restaurantID uint) error

	// HandleGetMenuItemsFromImage Wraper for handle GetMenuItemsFromImage function
	HandleGetMenuItemsFromImage(data string, restaurantID uint) error

	// AnalyzeImage gets the text from the image and returns the extracted menu items.
	AnalyzeImage(fileBytes []byte, restaurantID uint) (string, error)
}
