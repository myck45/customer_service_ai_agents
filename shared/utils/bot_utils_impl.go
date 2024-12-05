package utils

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/proyectos01-a/shared/data"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type BotUtilsImpl struct {
	openai   *openai.Client
	menuRepo data.MenuRepository
	botTools BotTools
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
				Type:     openai.ToolTypeFunction,
				Function: b.botTools.GetMenuItemsFromImage(),
			},
		},
		ToolChoice: openai.ToolChoice{
			Type: openai.ToolTypeFunction,
			Function: openai.ToolFunction{
				Name: "get_menu_items_from_image",
			},
		},
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Extract the menu items from the image",
			},
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeImageURL,
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

	if resp.Choices[0].FinishReason != openai.FinishReasonToolCalls {
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
