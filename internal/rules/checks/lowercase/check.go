package lowercase

import (
	"strings"
	"unicode"

	"example.com/slogzaplint/internal/rules/core"
	"example.com/slogzaplint/internal/rules/helpers"
)

func CheckLowercaseStart(msg string) []core.Violation {
	trimmed := strings.TrimSpace(msg)
	if trimmed == "" {
		return nil
	}

	if firstLetter, ok := helpers.FirstUnicodeLetter(trimmed); ok {
		if !unicode.IsLower(firstLetter) {
			return []core.Violation{{
				Kind:    core.KindLowercase,
				Message: "message must start with a lowercase letter",
			}}
		}
	}
	return nil
}
