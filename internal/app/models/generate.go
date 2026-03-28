package models

import "encoding/json"

type LLMType string

const (
	TypeChat  LLMType = "chat"
	TypeImage LLMType = "image"
	TypeAudio LLMType = "audio"
)

type GenerateRequest struct {
	Type    LLMType         `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type GenerateResponse struct {
	Result any `json:"result"`
}
