package models

type ImagePayload struct {
	Prompt string `json:"prompt"`
	Size   string `json:"size"`
}
