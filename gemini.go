package main

import (
	"context"
	"fmt"
)

type GeminiGateway struct {
	APIKey string
}

func NewGeminiGateway(apiKey string) *GeminiGateway {
	return &GeminiGateway{APIKey: apiKey}
}

func (g *GeminiGateway) Generate(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("Response from Gemini for: %s", prompt), nil
}
