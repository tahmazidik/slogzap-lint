package gclplugin

import (
	"fmt"
	"strings"

	"example.com/slogzaplint/internal/analyzer"
)

func decodeSettings(in any, base analyzer.Settings) (analyzer.Settings, error) {
	m, ok := in.(map[string]any)
	if !ok {
		return base, fmt.Errorf("settings must be a map, got %T", in)
	}

	out := base

	if v, ok := m["levels"]; ok {
		arr, err := toStringSliceFlexible(v)
		if err != nil {
			return base, fmt.Errorf("levels: %w", err)
		}
		out.Levels = normalizeLevels(arr)
	}

	if v, ok := m["sensitiveKeys"]; ok {
		arr, err := toStringSliceFlexible(v)
		if err != nil {
			return base, fmt.Errorf("sensitiveKeys: %w", err)
		}
		out.SensitiveKeys = cleanStrings(arr)
	}

	if v, ok := m["sensitive_keys"]; ok {
		arr, err := toStringSliceFlexible(v)
		if err != nil {
			return base, fmt.Errorf("sensitive_keys: %w", err)
		}
		out.SensitiveKeys = cleanStrings(arr)
	}

	return out, nil
}

func toStringSliceFlexible(v any) ([]string, error) {
	switch x := v.(type) {
	case string:
		return []string{strings.TrimSpace(x)}, nil
	case []string:
		out := make([]string, 0, len(x))
		for i, s := range x {
			if strings.TrimSpace(s) == "" {
				return nil, fmt.Errorf("item %d is empty", i)
			}
			out = append(out, strings.TrimSpace(s))
		}
		return out, nil
	case []any:
		out := make([]string, 0, len(x))
		for i, it := range x {
			s, ok := it.(string)
			if !ok {
				return nil, fmt.Errorf("item %d: expected string, got %T", i, it)
			}
			s = strings.TrimSpace(s)
			if s == "" {
				return nil, fmt.Errorf("item %d is empty", i)
			}
			out = append(out, s)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("expected string or array of strings, got %T", v)
	}
}

func normalizeLevels(in []string) []string {
	out := make([]string, 0, len(in))
	for _, l := range in {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}
		lower := strings.ToLower(l)
		switch lower {
		case "debug":
			out = append(out, "debug")
		case "info":
			out = append(out, "info")
		case "warn":
			out = append(out, "warn")
		case "error":
			out = append(out, "error")
		default:
			out = append(out, l)
		}
	}
	return out
}

func cleanStrings(in []string) []string {
	out := make([]string, 0, len(in))
	for _, s := range in {
		s = strings.TrimSpace(s)
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
