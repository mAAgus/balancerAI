package llmmodels

import (
	"balancerAI/internal/app/models"
	"context"
)

type ChatModel interface {
	GenerateChat(ctx context.Context, payload models.ChatPayload) (string, error)
}

type ImageModel interface {
	GenerateImage(ctx context.Context, payload models.ImagePayload) (string, error)
}
type AudioModel interface {
	GenerateAudio(ctx context.Context, payload models.AudioPayload) (string, error)
}
