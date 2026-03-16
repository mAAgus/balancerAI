package balancer

import (
	llmmodels "balancerAI/internal/app/llm_models"
	"balancerAI/internal/app/models"
	"context"
	"errors"
	"sync"
)

type Balancer struct {
	chatModels  []llmmodels.ChatModel
	imageModels []llmmodels.ImageModel
	audioModels []llmmodels.AudioModel
	chatIndex   int
	imageIndex  int
	audioIndex  int

	mu sync.Mutex
}

func NewBalancer(chatModels []llmmodels.ChatModel,
	imageModels []llmmodels.ImageModel,
	audioModels []llmmodels.AudioModel) *Balancer {
	return &Balancer{
		chatModels:  chatModels,
		imageModels: imageModels,
		audioModels: audioModels,
	}
}
func (b *Balancer) nextChatModel() llmmodels.ChatModel {
	b.mu.Lock()
	defer b.mu.Unlock()
	model := b.chatModels[b.chatIndex]
	b.chatIndex = (b.chatIndex + 1) % len(b.chatModels)
	return model
}
func (b *Balancer) nextImageModel() llmmodels.ImageModel {
	b.mu.Lock()
	defer b.mu.Unlock()
	model := b.imageModels[b.imageIndex]
	b.imageIndex = (b.imageIndex + 1) % len(b.imageModels)
	return model
}
func (b *Balancer) nextAudioModel() llmmodels.AudioModel {
	b.mu.Lock()
	defer b.mu.Unlock()
	model := b.audioModels[b.audioIndex]
	b.audioIndex = (b.audioIndex + 1) % len(b.audioModels)
	return model
}
func (b *Balancer) GenerateChat(ctx context.Context, payload models.ChatPayload) (string, error) {
	if len(b.chatModels) == 0 {
		return "", errors.New("no chat models configured")
	}
	var lastErr error
	for i := 0; i < len(b.chatModels); i++ {
		model := b.nextChatModel()
		resp, err := model.GenerateChat(ctx, payload)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return "", lastErr
}
func (b *Balancer) GenerateImage(ctx context.Context, payload models.ImagePayload) (string, error) {
	if len(b.imageModels) == 0 {
		return "", errors.New("no image models configured")
	}
	var lastErr error
	for i := 0; i < len(b.imageModels); i++ {
		model := b.nextImageModel()
		resp, err := model.GenerateImage(ctx, payload)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return "", lastErr
}
func (b *Balancer) GenerateAudio(ctx context.Context, payload models.AudioPayload) (string, error) {
	if len(b.audioModels) == 0 {
		return "", errors.New("no audio models configured")
	}
	var lastErr error
	for i := 0; i < len(b.audioModels); i++ {
		model := b.nextAudioModel()
		resp, err := model.GenerateAudio(ctx, payload)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return "", lastErr
}
