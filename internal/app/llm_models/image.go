package llmmodels

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"balancerAI/internal/app/models"
)

type ImageLLM struct {
	endpoint string
	apiKey   string
	client   *http.Client
}

func NewImageLLM(endpoint, apiKey string) *ImageLLM {
	return &ImageLLM{
		endpoint: endpoint,
		apiKey:   apiKey,
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

type imageRequest struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
}

type imageResponse struct {
	URL string `json:"url"`
}

func (i *ImageLLM) GenerateImage(ctx context.Context, payload models.ImagePayload) (string, error) {

	reqBody := imageRequest{
		Prompt: payload.Prompt,
		Size:   payload.Size,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		i.endpoint,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+i.apiKey)

	resp, err := i.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("image model returned non-200 status")
	}

	var result imageResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
