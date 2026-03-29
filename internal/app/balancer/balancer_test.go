package balancer_test

import (
	"balancerAI/internal/app/balancer"
	llmmodels "balancerAI/internal/app/llm_models"
	"balancerAI/internal/app/models"
	"context"
	"errors"
	"sync"
	"testing"
)

type mockChatModel struct {
	response string
	err      error
}

func (m *mockChatModel) GenerateChat(ctx context.Context, payload models.ChatPayload) (string, error) {
	return m.response, m.err
}
func TestGenerateChatSuccess(t *testing.T) {
	model := &mockChatModel{
		response: "hello",
	}
	b := balancer.NewBalancer([]llmmodels.ChatModel{model}, nil, nil)
	resp, err := b.GenerateChat(context.Background(), models.ChatPayload{Message: "hi"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "hello" {
		t.Fatalf("expected hello got %s", resp)
	}
}
func TestGenerateChatFallback(t *testing.T) {
	model1 := &mockChatModel{
		err: errors.New("model failed"),
	}
	model2 := &mockChatModel{
		response: "fallback response",
	}
	b := balancer.NewBalancer([]llmmodels.ChatModel{model1, model2}, nil, nil)
	resp, err := b.GenerateChat(context.Background(), models.ChatPayload{
		Message: "hi",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp != "fallback response" {
		t.Fatalf("expected fallback response got %s", resp)
	}
}
func TestGenerateChatNoModels(t *testing.T) {
	b := balancer.NewBalancer(nil, nil, nil)
	_, err := b.GenerateChat(context.Background(), models.ChatPayload{Message: "hi"})
	if err == nil {
		t.Fatalf("expexted error dut got nil")
	}
}

type mockImageModel struct {
	response string
	err      error
}

func (m *mockImageModel) GenerateImage(ctx context.Context, payload models.ImagePayload) (string, error) {
	return m.response, nil
}

func TestGenerateImageSuccess(t *testing.T) {

	model := &mockImageModel{
		response: "image-url",
	}

	b := balancer.NewBalancer(
		nil,
		[]llmmodels.ImageModel{model},
		nil,
	)

	resp, err := b.GenerateImage(context.Background(), models.ImagePayload{
		Prompt: "cat",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp != "image-url" {
		t.Fatalf("expected image-url got %s", resp)
	}
}

type mockAudioModel struct {
	response string
	err      error
}

func (m *mockAudioModel) GenerateAudio(ctx context.Context, payload models.AudioPayload) (string, error) {
	return m.response, nil
}

func TestGenerateAudioSuccess(t *testing.T) {

	model := &mockAudioModel{
		response: "audio-url",
	}

	b := balancer.NewBalancer(
		nil,
		nil,
		[]llmmodels.AudioModel{model},
	)

	resp, err := b.GenerateAudio(context.Background(), models.AudioPayload{
		Text: "hello",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp != "audio-url" {
		t.Fatalf("expected audio-url got %s", resp)
	}
}
func TestRoundRobinChat(t *testing.T) {

	model1 := &mockChatModel{response: "1"}
	model2 := &mockChatModel{response: "2"}

	b := balancer.NewBalancer(
		[]llmmodels.ChatModel{model1, model2},
		nil,
		nil,
	)

	r1, _ := b.GenerateChat(context.Background(), models.ChatPayload{})
	r2, _ := b.GenerateChat(context.Background(), models.ChatPayload{})
	r3, _ := b.GenerateChat(context.Background(), models.ChatPayload{})

	if r1 != "1" {
		t.Fatalf("expected 1 got %s", r1)
	}

	if r2 != "2" {
		t.Fatalf("expected 2 got %s", r2)
	}

	if r3 != "1" {
		t.Fatalf("expected 1 got %s", r3)
	}
}
func TestGenerateChatConcurrency(t *testing.T) {

	model := &mockChatModel{
		response: "ok",
	}

	b := balancer.NewBalancer(
		[]llmmodels.ChatModel{model},
		nil,
		nil,
	)

	const workers = 100

	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {

		go func() {
			defer wg.Done()

			resp, err := b.GenerateChat(
				context.Background(),
				models.ChatPayload{Message: "hi"},
			)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if resp != "ok" {
				t.Errorf("unexpected response: %s", resp)
			}
		}()

	}

	wg.Wait()
}
func TestRoundRobinConcurrency(t *testing.T) {

	model1 := &mockChatModel{response: "1"}
	model2 := &mockChatModel{response: "2"}

	b := balancer.NewBalancer(
		[]llmmodels.ChatModel{model1, model2},
		nil,
		nil,
	)

	const workers = 100

	var wg sync.WaitGroup
	wg.Add(workers)

	results := make(chan string, workers)

	for i := 0; i < workers; i++ {

		go func() {
			defer wg.Done()

			resp, _ := b.GenerateChat(
				context.Background(),
				models.ChatPayload{},
			)

			results <- resp
		}()

	}

	wg.Wait()
	close(results)

	count1 := 0
	count2 := 0

	for r := range results {
		if r == "1" {
			count1++
		}
		if r == "2" {
			count2++
		}
	}

	if count1 == 0 || count2 == 0 {
		t.Fatalf("round robin failed: model1=%d model2=%d", count1, count2)
	}
}
