package service_test

import (
	"balancerAI/internal/app/models"
	"balancerAI/internal/app/service"
	"context"
	"encoding/json"
	"testing"
)

type mockBalalancer struct{}

func (m *mockBalalancer) GenerateChat(cyx context.Context, payload models.ChatPayload) (string, error) {
	return "chat-result", nil
}
func (m *mockBalalancer) GenerateImage(cyx context.Context, payload models.ImagePayload) (string, error) {
	return "image-result", nil
}
func (m *mockBalalancer) GenerateAudio(cyx context.Context, payload models.AudioPayload) (string, error) {
	return "audio-result", nil
}

func TestServiceGenerateChat(t *testing.T) {
	b := &mockBalalancer{}
	s := service.NewService(b)
	payload := models.ChatPayload{
		Message: "hello",
	}
	raw, _ := json.Marshal(payload)

	req := models.GenerateRequest{
		Type:    models.TypeChat,
		Payload: raw,
	}
	resp, err := s.Generate(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "chat-result" {
		t.Fatalf("expected chat-result got %v", resp)
	}
}
func TestServiceGenerateImage(t *testing.T) {
	b := &mockBalalancer{}
	s := service.NewService(b)
	payload := models.ImagePayload{
		Prompt: "cat",
		Size:   "1024x1024",
	}
	raw, _ := json.Marshal(payload)

	req := models.GenerateRequest{
		Type:    models.TypeImage,
		Payload: raw,
	}
	resp, err := s.Generate(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "image-result" {
		t.Fatalf("expected image-result got %v", resp)
	}
}
func TestServiceGenerateAudio(t *testing.T) {
	b := &mockBalalancer{}
	s := service.NewService(b)
	payload := models.AudioPayload{
		Text:  "hello",
		Voice: "male",
	}
	raw, _ := json.Marshal(payload)

	req := models.GenerateRequest{
		Type:    models.TypeAudio,
		Payload: raw,
	}
	resp, err := s.Generate(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "audio-result" {
		t.Fatalf("expected audio-result got %v", resp)
	}
}
func TestServiceUnknownType(t *testing.T) {
	b := &mockBalalancer{}
	s := service.NewService(b)
	req := models.GenerateRequest{
		Type: "unknown",
	}
	_, err := s.Generate(context.Background(), req)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
}
