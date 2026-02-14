package a

import (
	"log/slog"
)

func SlogCases() {
	slog.Info("Hello world")
	slog.Info("привет")
	slog.Info("hello!")
	slog.Info("token expired")
	slog.Info("hello world")
}

func ZapCases() {
	zap.L().Info("Hello")
	zap.L().Info("token ok")
	zap.L().Info("hello!")
	zap.L().Info("hello world")
}
