package api

import (
	"balancerAI/internal/app/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type mockService struct{}

func (m *mockService) Generate(ctx context.Context, req models.GenerateRequest) (any, error) {
	return "mock-result", nil
}
func TestGenerateHandlerSuccess(t *testing.T) {
	app := fiber.New()
	api := &API{
		app:     app,
		service: &mockService{},
		logger:  logrus.New(),
	}
	app.Post("/generate", api.handleGenerate)
	payload := models.GenerateRequest{
		Type:    models.TypeChat,
		Payload: json.RawMessage(`{"message":"hello"}`),
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
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200 got %d", resp.StatusCode)
	}
}
func TestGenerateHandlerInvalidJSON(t *testing.T) {

	app := fiber.New()

	api := &API{
		app:     app,
		service: &mockService{},
		logger:  logrus.New(),
	}

	app.Post("/generate", api.handleGenerate)

	req := httptest.NewRequest(
		"POST",
		"/generate",
		bytes.NewReader([]byte("invalid-json")),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("expected 400 got %d", resp.StatusCode)
	}
}

type mockServiceError struct{}

func (m *mockServiceError) Generate(ctx context.Context, req models.GenerateRequest) (any, error) {
	return nil, errors.New("generation failed")
}
func TestGenerateHandlerServiceError(t *testing.T) {

	app := fiber.New()

	api := &API{
		app:     app,
		service: &mockServiceError{},
		logger:  logrus.New(),
	}

	app.Post("/generate", api.handleGenerate)

	payload := models.GenerateRequest{
		Type:    models.TypeChat,
		Payload: json.RawMessage(`{"message":"hello"}`),
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

	if resp.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("expected %d got %d", fiber.StatusInternalServerError, resp.StatusCode)
	}
}
func TestHealthHandler(t *testing.T) {

	app := fiber.New()

	api := &API{
		app:    app,
		logger: logrus.New(),
	}

	app.Get("/health", api.handleHealth)

	req := httptest.NewRequest(
		"GET",
		"/health",
		nil,
	)

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected %d got %d", fiber.StatusOK, resp.StatusCode)
	}
}
