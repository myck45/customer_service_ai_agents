package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/pgvector/pgvector-go"
	"github.com/proyectos01-a/shared/data"
	"github.com/proyectos01-a/shared/dto/res"
	"github.com/proyectos01-a/shared/models"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type BotUtilsImpl struct {
	openai   *openai.Client
	menuRepo data.MenuRepository
	botTools BotTools
}

// HandleGetUserOrder implements BotUtils.
func (b *BotUtilsImpl) HandleGetUserOrder(data string, restaurantID uint) (string, error) {
	panic("unimplemented")
}

// HandleGetMenuItemsFromImage implements BotUtils.
func (b *BotUtilsImpl) HandleGetMenuItemsFromImage(data string, restaurantID uint) (string, error) {

	// Parse the data into a slice of ExtractedMenuItemResponse
	var extractedItems []res.ExtractedMenuItemResponse
	if err := json.Unmarshal([]byte(data), &extractedItems); err != nil {
		logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to unmarshal data")
		return "", err
	}

	// Iterate over the extracted items and create a menu for each
	for _, item := range extractedItems {
		// Marshal the item into a string to generate an embedding
		itemStr, err := json.Marshal(item)
		if err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to marshal item")
			return "", err
		}
		embedding, err := b.GenerateEmbedding(string(itemStr))
		if err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to generate embedding")
			return "", err
		}

		// Create a new menu by each item
		menu := &models.Menu{
			ItemName:     item.ItemName,
			Description:  item.Description,
			Price:        item.Price,
			Likes:        0,
			Embedding:    pgvector.NewVector(embedding),
			RestaurantID: restaurantID,
		}
		if err := b.menuRepo.CreateMenu(menu); err != nil {
			logrus.WithError(err).Error("[HandleGetMenuItemsFromImage] failed to create menu")
			return "", err
		}
	}

	return "success", nil
}

// BotToolsHandler implements BotUtils.
func (b *BotUtilsImpl) BotToolsHandler(functionName string, data string, restaurantID uint) (string, error) {
	switch functionName {
	case "get_user_order":
		return "get_user_order", nil
	case "extract_menu_items":
		return b.HandleGetMenuItemsFromImage(data, restaurantID)
	default:
		return "", errors.New("function not found")
	}
}

// GenerateEmbedding implements BotUtils.
func (b *BotUtilsImpl) GenerateEmbedding(data string) ([]float32, error) {
	if data == "" {
		return nil, errors.New("input data cannot be empty")
	}

	ctx := context.Background()

	targetReq := openai.EmbeddingRequest{
		Input: []string{data},
		Model: openai.LargeEmbedding3,
	}

	res, err := b.openai.CreateEmbeddings(ctx, targetReq)
	if err != nil {
		logrus.WithError(err).Error("failed to create embeddings")
		return nil, err
	}

	embedding := res.Data[0].Embedding

	return embedding, nil
}

// AnalyzeImage implements BotUtils.
func (b *BotUtilsImpl) AnalyzeImage(fileBytes []byte, restaurantID uint) (string, error) {

	ctx := context.Background()

	// Encode the image to base64
	imageBase64 := base64.StdEncoding.EncodeToString(fileBytes)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Tools: []openai.Tool{
			{
				Type:     "function",
				Function: b.botTools.getMenuItemsFromImage(),
			},
		},
		ToolChoice: openai.ToolChoice{
			Type: "function",
			Function: openai.ToolFunction{
				Name: "extract_menu_items",
			},
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: "Extract the menu items from the image",
			},
			{
				Role: "user",
				MultiContent: []openai.ChatMessagePart{
					{
						Type: "image_utl",
						ImageURL: &openai.ChatMessageImageURL{
							URL: fmt.Sprintf("data:image/jpeg;base64,%s", imageBase64),
						},
					},
				},
			},
		},
	}

	resp, err := b.openai.CreateChatCompletion(ctx, req)
	if err != nil {
		logrus.WithError(err).Error("[AnalyzeImage] failed to analyze image")
		return "", err
	}

	if resp.Choices[1].FinishReason != openai.FinishReasonToolCalls {
		logrus.Error("[AnalyzeImage] tool call did not finish")
		return "", errors.New("tool call did not finish")
	}

	// Handle the response
	if len(resp.Choices[0].Message.ToolCalls) == 0 {
		logrus.Error("[AnalyzeImage] no tool call returned")
		return "", errors.New("no menu items extracted")
	}

	// Parse the tool call arguments
	toolCall := resp.Choices[0].Message.ToolCalls[0]

	var functionArgs map[string]interface{}
	err = json.Unmarshal([]byte(toolCall.Function.Arguments), &functionArgs)
	if err != nil {
		logrus.WithError(err).Error("[AnalyzeImage] failed to unmarshal tool call arguments")
		return "", err
	}

	// Convert function arguments to JSON string for HandleGetMenuItemsFromImage
	jsonData, err := json.Marshal(functionArgs["items"])
	if err != nil {
		logrus.WithError(err).Error("[AnalyzeImage] failed to marshal items")
		return "", err
	}

	// Return the JSON string
	return string(jsonData), nil
}

func NewBotUtilsImpl(openai *openai.Client, menuRepo data.MenuRepository, botTools BotTools) BotUtils {
	return &BotUtilsImpl{
		openai:   openai,
		menuRepo: menuRepo,
		botTools: botTools,
	}
}
