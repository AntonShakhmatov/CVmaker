package gemini

import "fmt"

func buildPrompt(text string) string {
	return fmt.Sprintf(`V této zprávě je Portfolio. Tvůj úkol: udělat zavěr po informace v tomto cv.%s`, text)
}
