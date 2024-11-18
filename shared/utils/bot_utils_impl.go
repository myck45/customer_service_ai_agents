package utils

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type BotUtilsImpl struct {
	openai *openai.Client
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

func NewBotUtilsImpl(openai *openai.Client) BotUtils {
	return &BotUtilsImpl{openai: openai}
}
