package rules

import (
	"strconv"
	"strings"
	"unicode"
)

var sensitiveKeywords = []string{
	"password", "passwd", "pwd",
	"token", "access_token", "refresh_token",
	"secret", "api_key", "apikey",
	"authorization", "bearer",
	"cookie", "session",
}

func ValidateMessage(msg string) []string {
	var out []string
	trimmed := strings.TrimSpace(msg)
	if trimmed == "" {
		return out
	}

	if firstLetter, ok := firstUnicodeLetter(trimmed); ok {
		if !unicode.IsLower(firstLetter) {
			out = append(out, "message must start with a lowercase letter")
		}
	}

	if containsCyrillic(trimmed) {
		out = append(out, "message must be English-only (Cyrillic detected)")
	}

	if bad, ok := firstForbiddenSymbol(trimmed); ok {
		out = append(out, "message contains forbidden symbol: "+strconv.QuoteRune(bad))
	}

	if kw, ok := containsSensitiveKeyword(trimmed); ok {
		out = append(out, "message my contain sensitive data (keyword: "+kw+")")
	}
	return out
}

func firstUnicodeLetter(s string) (rune, bool) {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return r, true
		}
	}
	return 0, false
}

func containsCyrillic(s string) bool {
	for _, r := range s {
		if r >= 0x0400 && r <= 0x04FF {
			return true
		}
	}
	return false
}

func firstForbiddenSymbol(s string) (rune, bool) {
	for _, r := range s {
		if r == ' ' {
			continue
		}

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			continue
		}

		return r, true
	}
	return 0, false
}

func containsSensitiveKeyword(s string) (string, bool) {
	low := strings.ToLower(s)
	for _, kw := range sensitiveKeywords {
		if strings.Contains(low, kw) {
			return kw, true
		}
	}
	return "", false
}
