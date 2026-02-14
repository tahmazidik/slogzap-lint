package a

import (
	"log/slog"
)

func SlogCases() {
	slog.Info("Hello world")   // want `slog info: message must start with a lowercase letter`
	slog.Info("привет")        // want `slog info: message must be English-only \(Cyrillic detected\)`
	slog.Info("hello!")        // want `slog info: message contains forbidden symbol: '!'`
	slog.Info("token expired") // want `slog info: message may contain sensitive data \(keyword: token\)`
	slog.Info("hello world")
}

func ZapCases() {
	zap.L().Info("Hello")    // want `zap info: message must start with a lowercase letter`
	zap.L().Info("token ok") // want `zap info: message may contain sensitive data \(keyword: token\)`
	zap.L().Info("hello!")   // want `zap info: message contains forbidden symbol: '!'`
	zap.L().Info("hello world")
}
