package llmmodels

import (
	"balancerAI/internal/app/models"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ChatLLM struct {
	endpoint string
	apiKey   string
	client   *http.Client
}

func NewChatLLM(endpoint, apiKey string) *ChatLLM {
	return &ChatLLM{
		endpoint: endpoint,
		apiKey:   apiKey,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type externalChatRequest struct {
	Message string `json:"message"`
}
type externalChatResponse struct {
	Response string `json:"response"`
}

func (c *ChatLLM) GenerateChat(ctx context.Context, payload models.ChatPayload) (string, error) {
	reqBody := externalChatRequest{
		Message: payload.Message,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.endpoint,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("llm error: %s", string(body))
	}
	var result externalChatResponse
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Response, nil
}
