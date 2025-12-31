package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"backend/req/http/contacts"
	"backend/req/http/experience"
	"backend/req/http/header"
	"backend/req/http/language"
	"backend/req/http/links"
	"backend/req/http/projects"
	"backend/req/http/responsibilities"
	"backend/req/http/technologies"
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

func (c *GeminiClient) MakeLanguage(ctx context.Context, cvText string) (*language.LanguageForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out language.LanguageForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultLanguage(&out)
	return &out, nil
}

func trimResultLanguage(out *language.LanguageForm) {
	out.Language = strings.TrimSpace(out.Language)
}

func (c *GeminiClient) MakeHeader(ctx context.Context, cvText string) (*header.HeaderForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out header.HeaderForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultHeader(&out)
	return &out, nil
}

func trimResultHeader(out *header.HeaderForm) {
	out.FirstName = strings.TrimSpace(out.FirstName)
	out.LastName = strings.TrimSpace(out.LastName)
	out.Competence = strings.TrimSpace(out.Competence)
	out.Title = strings.TrimSpace(out.Title)
}

func (c *GeminiClient) MakeContacts(ctx context.Context, cvText string) (*contacts.ContactsForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out contacts.ContactsForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultContacts(&out)
	return &out, nil
}

func trimResultContacts(out *contacts.ContactsForm) {
	out.Phone = strings.TrimSpace(out.Phone)
	out.Email = strings.TrimSpace(out.Email)
	out.Address = strings.TrimSpace(out.Address)
}

func (c *GeminiClient) MakeExperience(ctx context.Context, cvText string) (*experience.ExperienceForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out experience.ExperienceForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultExperience(&out)
	return &out, nil
}

func trimResultExperience(out *experience.ExperienceForm) {
	out.Company = strings.TrimSpace(out.Company)
	out.Location_field = strings.TrimSpace(out.Location_field)
	out.Position = strings.TrimSpace(out.Position)
	out.Date_from = strings.TrimSpace(out.Date_from)
	out.Date_to = strings.TrimSpace(out.Date_to)
}

func (c *GeminiClient) MakeProjects(ctx context.Context, cvText string) (*projects.ProjectsForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out projects.ProjectsForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultProjects(&out)
	return &out, nil
}

func trimResultProjects(out *projects.ProjectsForm) {
	out.Name = strings.TrimSpace(out.Name)
	out.Description = strings.TrimSpace(out.Description)
}

func (c *GeminiClient) MakeResponsibilities(ctx context.Context, cvText string) (*responsibilities.ResponsibilitiesForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out responsibilities.ResponsibilitiesForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultResponsibilities(&out)
	return &out, nil
}

func trimResultResponsibilities(out *responsibilities.ResponsibilitiesForm) {
	out.Description = strings.TrimSpace(out.Description)
}

func (c *GeminiClient) MakeTechnologies(ctx context.Context, cvText string) (*technologies.TechnologiesForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out technologies.TechnologiesForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultTechnologies(&out)
	return &out, nil
}

func trimResultTechnologies(out *technologies.TechnologiesForm) {
	out.Name = strings.TrimSpace(out.Name)
}

func (c *GeminiClient) MakeLinks(ctx context.Context, cvText string) (*links.LinksForm, error) {
	prompt := buildPrompt(cvText)

	reqBody := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]any{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(reqBody)

	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s",
		c.model, c.apiKey,
	)

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("gemini http %d: %s", resp.StatusCode, string(raw))
	}

	modelText := extractModelText(raw)
	jsonStr, err := extractJSON(modelText)
	if err != nil {
		return nil, err
	}

	var out links.LinksForm
	if err := json.Unmarshal([]byte(jsonStr), &out); err != nil {
		return nil, err
	}

	trimResultLinks(&out)
	return &out, nil
}

func trimResultLinks(out *links.LinksForm) {
	out.Label = strings.TrimSpace(out.Label)
}
