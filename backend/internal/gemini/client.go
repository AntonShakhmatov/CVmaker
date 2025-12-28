package gemini

import (
	"net/http"
	"time"
)

type GeminiClient struct {
	apiKey     string
	model      string
	httpClient *http.Client
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey: apiKey,
		model:  "gemini-2.0-flash",
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// func (c *GeminiClient) MakeHeader(ctx context.Context, wpsText string) (*ResultHeader, error) {
// 	prompt := buildPrompt(wpsText)

// 	reqBody := map[string]any{
// 		"contents": []any{
// 			map[string]any{
// 				"parts": []any{
// 					map[string]any{"text": prompt},
// 				},
// 			},
// 		},
// 	}

// 	b, _ := json.Marshal(reqBody)

// 	url := fmt.Sprintf(
// 		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
// 		c.model, c.apiKey,
// 	)

// 	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// }
