package symbols

import (
	"strconv"
	"strings"

	"example.com/slogzaplint/internal/rules/core"
	"example.com/slogzaplint/internal/rules/helpers"
)

func CheckForbiddenSymbols(msg string) []core.Violation {
	trimmed := strings.TrimSpace(msg)
	if trimmed == "" {
		return nil
	}

	if bad, ok := helpers.FirstForbiddenSymbol(trimmed); ok {
		return []core.Violation{{
			Kind:    core.KindForbidden,
			Message: "message contains forbidden symbol: " + strconv.QuoteRune(bad),
			BadRune: bad,
		}}
	}
	return nil
}
