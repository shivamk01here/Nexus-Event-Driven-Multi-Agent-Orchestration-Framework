package main

import (
	"context"
	"fmt"
)

type OpenAIGateway struct {
	APIKey string
}

func NewOpenAIGateway(apiKey string) *OpenAIGateway {
	return &OpenAIGateway{APIKey: apiKey}
}

func (g *OpenAIGateway) Generate(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("Response from OpenAI for: %s", prompt), nil
}
