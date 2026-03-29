package service

import (
	"balancerAI/internal/app/models"
	"context"
	"encoding/json"
	"errors"

	"go.opentelemetry.io/otel"
)

type Balancer interface {
	GenerateChat(ctx context.Context, payload models.ChatPayload) (string, error)
	GenerateImage(ctx context.Context, payload models.ImagePayload) (string, error)
	GenerateAudio(ctx context.Context, payload models.AudioPayload) (string, error)
}
type Service struct {
	balancer Balancer
}

func NewService(b Balancer) *Service {
	return &Service{
		balancer: b,
	}
}

func (s *Service) Generate(ctx context.Context, req models.GenerateRequest) (any, error) {
	tracer := otel.Tracer("ai-balancer")

	ctx, span := tracer.Start(ctx, "service.Generate")
	defer span.End()

	switch req.Type {

	case models.TypeChat:

		var payload models.ChatPayload

		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			return nil, err
		}

		return s.balancer.GenerateChat(ctx, payload)

	case models.TypeAudio:

		var payload models.AudioPayload

		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			return nil, err
		}

		return s.balancer.GenerateAudio(ctx, payload)

	case models.TypeImage:

		var payload models.ImagePayload

		if err := json.Unmarshal(req.Payload, &payload); err != nil {
			return nil, err
		}

		return s.balancer.GenerateImage(ctx, payload)

	default:
		return nil, errors.New("unknown request type")
	}

}
