package rules

import (
	"example.com/slogzaplint/internal/rules/checks/language"
	"example.com/slogzaplint/internal/rules/checks/lowercase"
	"example.com/slogzaplint/internal/rules/checks/sensitive"
	"example.com/slogzaplint/internal/rules/checks/symbols"
	"example.com/slogzaplint/internal/rules/core"
)

type Violation = core.Violation
type Kind = core.Kind

func ValidateMessage(msg string, sensitiveKeys []string) []core.Violation {
	var out []core.Violation

	out = append(out, lowercase.CheckLowercaseStart(msg)...)

	cyr := language.CheckEnglishOnly(msg)
	out = append(out, cyr...)
	if len(cyr) == 0 {
		out = append(out, symbols.CheckForbiddenSymbols(msg)...)
	}

	out = append(out, sensitive.CheckSensitive(msg, sensitiveKeys)...)
	return out
}
