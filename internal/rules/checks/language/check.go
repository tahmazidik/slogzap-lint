package language

import (
	"strings"

	"example.com/slogzaplint/internal/rules/core"
	"example.com/slogzaplint/internal/rules/helpers"
)

func CheckEnglishOnly(msg string) []core.Violation {
	trimmed := strings.TrimSpace(msg)
	if trimmed == "" {
		return nil
	}

	if helpers.ContainsCyrillic(trimmed) {
		return []core.Violation{{
			Kind:    core.KindCyrillic,
			Message: "message must be English-only (Cyrillic detected)",
		}}
	}
	return nil
}
