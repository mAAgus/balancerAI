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

type AudioLLM struct {
	endpoint string
	apiKey   string
	client   *http.Client
}

func NewAudioLLM(endpoint, apiKey string) *AudioLLM {
	return &AudioLLM{
		endpoint: endpoint,
		apiKey:   apiKey,
		client: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
}

type audioRequest struct {
	Text  string `json:"text"`
	Voice string `json:"voice"`
}

type audioResponse struct {
	URL string `json:"url"`
}

func (a *AudioLLM) GenerateAudio(ctx context.Context, payload models.AudioPayload) (string, error) {

	reqBody := audioRequest{
		Text:  payload.Text,
		Voice: payload.Voice,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		a.endpoint,
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.apiKey)

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("audio model returned non-200 status")
	}

	var result audioResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
