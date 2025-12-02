package logger

import (
	"L2.18/internal/config"
	"L2.18/pkg/logger/slog"
)

type Logger interface {
	LogFatal(msg string, err error, args ...any)
	LogError(msg string, err error, args ...any)
	LogWarn(msg string, args ...any)
	LogInfo(msg string, args ...any)
	Debug(msg string, args ...any)
	Close() error
}

func NewLogger(config config.Logger) Logger {
	return slog.NewLogger(config)
}
