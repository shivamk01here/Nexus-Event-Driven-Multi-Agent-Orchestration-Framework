package main

import (
	"context"
	"errors"
)

// LLMGateway provides a unified interface for multiple LLM providers.
type LLMGateway interface {
	Generate(ctx context.Context, prompt string) (string, error)
}

// GatewayProvider identifies the backend LLM provider.
type GatewayProvider string

const (
	ProviderOpenAI    GatewayProvider = "openai"
	ProviderAnthropic GatewayProvider = "anthropic"
	ProviderGemini    GatewayProvider = "gemini"
)

// NewLLMGateway acts as a factory for LLM gateways.
func NewLLMGateway(provider GatewayProvider, apiKey string) (LLMGateway, error) {
	switch provider {
	case ProviderOpenAI:
		return NewOpenAIGateway(apiKey), nil
	case ProviderAnthropic:
		return NewAnthropicGateway(apiKey), nil
	case ProviderGemini:
		return NewGeminiGateway(apiKey), nil
	default:
		return nil, errors.New("unsupported gateway provider")
	}
}
