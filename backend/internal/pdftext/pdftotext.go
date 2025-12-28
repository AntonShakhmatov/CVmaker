package pdftext

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func ToText(ctx context.Context, pdfPath string) (string, error) {
	cmd := exec.CommandContext(ctx, "pdftotext", "-layout", "-enc", "UTF-8", pdfPath, "-")

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pdftotext: %w; out=%s", err, out.String())
	}

	txt := strings.TrimSpace(strings.ReplaceAll(out.String(), "\u0000", ""))
	return txt, nil
}
