package gclplugin

import (
	"fmt"

	"example.com/slogzaplint/internal/analyzer"
)

func decodeSettings(in any, base analyzer.Settings) (analyzer.Settings, error) {
	m, ok := in.(map[string]any)
	if !ok {
		return base, fmt.Errorf("settings must be a map, got %T", in)
	}

	out := base

	if v, ok := m["levels"]; ok {
		arr, err := toStringSlice(v)
		if err != nil {
			return base, fmt.Errorf("levels: %w", err)
		}
		out.Levels = arr
	}

	if v, ok := m["sensitiveKeys"]; ok {
		arr, err := toStringSlice(v)
		if err != nil {
			return base, fmt.Errorf("sensitiveKeys: %w", err)
		}
		out.SensitiveKeys = arr
	}

	return out, nil
}

func toStringSlice(v any) ([]string, error) {
	switch x := v.(type) {
	case []string:
		return x, nil
	case []any:
		out := make([]string, 0, len(x))
		for _, it := range x {
			s, ok := it.(string)
			if !ok {
				return nil, fmt.Errorf("expected string item, got %T", it)
			}
			out = append(out, s)
		}
		return out, nil
	default:
		return nil, fmt.Errorf("expected array, got %T", v)
	}
}
