package main

import (
	"context"
	"fmt"
)

type AnthropicGateway struct {
	APIKey string
}

func NewAnthropicGateway(apiKey string) *AnthropicGateway {
	return &AnthropicGateway{APIKey: apiKey}
}

func (g *AnthropicGateway) Generate(ctx context.Context, prompt string) (string, error) {
	return fmt.Sprintf("Response from Anthropic for: %s", prompt), nil
}
