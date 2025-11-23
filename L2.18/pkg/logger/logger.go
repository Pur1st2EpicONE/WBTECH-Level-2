package logger

import (
	"L2.18/internal/config"
	"L2.18/pkg/logger/slog"
)

type Logger interface {
	LogFatal(msg string, err error, args ...any)
	LogError(string, error, ...any)
	LogInfo(msg string, args ...any)
	Debug(msg string, args ...any)
	Close() error
}

func NewLogger(config config.Logger) Logger {
	return slog.NewLogger(config)
}
