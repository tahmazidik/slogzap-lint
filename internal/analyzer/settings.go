package analyzer

import "fmt"

type Settings struct {
	Levels        []string `json:"levels" yaml:"levels"`
	SensitiveKeys []string `json:"sensitiveKeys" yaml:"sensitiveKeys"`
}

func DefaultSettings() Settings {
	return Settings{
		Levels:        []string{"Debug", "Info", "Warn", "Error"},
		SensitiveKeys: []string{"password", "pass", "pwd", "token", "api_key", "apikey", "secret", "key"},
	}
}

func (s Settings) Validate() error {
	if len(s.Levels) == 0 {
		return fmt.Errorf("levels must not be empty")
	}
	allowed := map[string]bool{"Debug": true, "Info": true, "Warn": true, "Error": true}
	for _, l := range s.Levels {
		if !allowed[l] {
			return fmt.Errorf("unknown level %q (allowed: Debug, Info, Warn, Error)", l)
		}
	}
	return nil
}
