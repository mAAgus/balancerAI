package api

import (
	"balancerAI/internal/app/balancer"
	llmmodels "balancerAI/internal/app/llm_models"
	"balancerAI/internal/app/models"
	"balancerAI/internal/app/service"

	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type mockChatModel struct{}

func (m *mockChatModel) GenerateChat(
	ctx context.Context,
	payload models.ChatPayload,
) (string, error) {
	return "integration-response", nil
}

func TestGenerateIntegration(t *testing.T) {

	model := &mockChatModel{}

	b := balancer.NewBalancer(
		[]llmmodels.ChatModel{model},
		nil,
		nil,
	)

	s := service.NewService(b)

	app := fiber.New()

	apiInstance := &API{
		app:     app,
		service: s,
		logger:  logrus.New(),
	}

	app.Post("/generate", apiInstance.handleGenerate)

	payload := models.GenerateRequest{
		Type: models.TypeChat,
		Payload: json.RawMessage(`{
			"message": "hello"
		}`),
	}

	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(
		"POST",
		"/generate",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected %d got %d", fiber.StatusOK, resp.StatusCode)
	}
}
