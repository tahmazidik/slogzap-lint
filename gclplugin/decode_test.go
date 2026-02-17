package gclplugin

import (
	"reflect"
	"testing"

	"example.com/slogzaplint/internal/analyzer"
)

func TestDecodeSettings_LevelsAsString_Normalizes(t *testing.T) {
	base := analyzer.DefaultSettings()

	in := map[string]any{
		"levels": "info",
	}

	out, err := decodeSettings(in, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"Info"}
	if !reflect.DeepEqual(out.Levels, want) {
		t.Fatalf("levels mismatch: got %#v, want %#v", out.Levels, want)
	}

	if !reflect.DeepEqual(out.SensitiveKeys, base.SensitiveKeys) {
		t.Fatalf("sensitive keys mismatch: got %#v, want %#v", out.SensitiveKeys, base.SensitiveKeys)
	}
}

func TestDecodeSettings_LevelsAsArrayAny_Normalizes(t *testing.T) {
	base := analyzer.DefaultSettings()

	in := map[string]any{
		"levels": []any{"DEBUG", "warn", "Error"},
	}

	out, err := decodeSettings(in, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"Debug", "Warn", "Error"}
	if !reflect.DeepEqual(out.Levels, want) {
		t.Fatalf("levels mismatch: got %#v, want %#v", out.Levels, want)
	}
}

func TestDecodeSettings_SensitiveKeysSnakeCaseAlias(t *testing.T) {
	base := analyzer.DefaultSettings()

	in := map[string]any{
		"sensitive_keys": []any{" token ", "password"},
	}

	out, err := decodeSettings(in, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"token", "password"}
	if !reflect.DeepEqual(out.SensitiveKeys, want) {
		t.Fatalf("sensitive keys mismatch: got %#v, want %#v", out.SensitiveKeys, want)
	}
}

func TestDecodeSettings_InvalidItemType_GivesHelpfulError(t *testing.T) {
	base := analyzer.DefaultSettings()

	in := map[string]any{
		"levels": []any{"Info", 123},
	}

	_, err := decodeSettings(in, base)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if got := err.Error(); got == "" {
		t.Fatalf("expected non-empty error")
	}
}
