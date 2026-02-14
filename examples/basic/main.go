package demo

import (
	"go.uber.org/zap"
	"log/slog"
)

func Example() {
	slog.Info("Hello world")
	slog.Info("привет")
	slog.Info("hello!")
	slog.Info("token expired")
	slog.Info("hello world")
}

func ExampleZap() {
	logger := zap.NewExample()
	logger.Info("Hello")
	logger.Info("token ok")
	zap.L().Info("hello!")
}
