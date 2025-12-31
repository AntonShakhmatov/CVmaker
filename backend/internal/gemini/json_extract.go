package gemini

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

func extractModelText(raw []byte) string {
	var resp map[string]any
	_ = json.Unmarshal(raw, &resp)

	cands, ok := resp["candidates"].([]any)
	if !ok || len(cands) == 0 {
		return ""
	}
	c0, _ := cands[0].(map[string]any)
	content, _ := c0["content"].(map[string]any)
	parts, _ := content["parts"].([]any)
	if len(parts) == 0 {
		return ""
	}
	p0, _ := parts[0].(map[string]any)
	t, _ := p0["text"].(string)
	return t
}

func extractJSON(s string) (string, error) {
	clean := strings.TrimSpace(s)

	re1 := regexp.MustCompile("(?is)```json\\s*(\\{.*?\\})\\s*```")
	if m := re1.FindStringSubmatch(clean); len(m) == 2 {
		return strings.TrimSpace(m[1]), nil
	}
	re2 := regexp.MustCompile("(?is)```\\s*(\\{.*?\\})\\s*```")
	if m := re2.FindStringSubmatch(clean); len(m) == 2 {
		return strings.TrimSpace(m[1]), nil
	}

	start := strings.IndexByte(clean, '{')
	end := strings.LastIndexByte(clean, '}')
	if start >= 0 && end > start {
		return strings.TrimSpace(clean[start : end+1]), nil
	}

	return "", errors.New("json not found in model output")
}
