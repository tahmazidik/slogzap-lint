package sensitive

import (
	"strings"

	"example.com/slogzaplint/internal/rules/core"
	"example.com/slogzaplint/internal/rules/helpers"
)

func CheckSensitive(msg string, keys []string) []core.Violation {
	trimmed := strings.TrimSpace(msg)
	if trimmed == "" {
		return nil
	}

	low := strings.ToLower(trimmed)

	list := keys
	if len(list) == 0 {
		list = helpers.DefaultSensitiveKeyWords
	}

	for _, kw := range list {
		if kw == "" {
			continue
		}
		if strings.Contains(low, strings.ToLower(kw)) {
			return []core.Violation{{
				Kind:    core.KindSensitive,
				Message: "message may contain sensitive data (keyword: " + kw + ")",
				Keyword: kw,
			}}
		}
	}
	return nil
}
