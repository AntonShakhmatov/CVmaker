package cv

import (
	"backend/internal/gemini"
)

type Service struct {
	gemini *gemini.GeminiClient
}

func NewService(g *gemini.GeminiClient) *Service {
	return &Service{gemini: g}
}
