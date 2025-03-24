package common

import (
	"log/slog"
	"os"
)

var logLevel = new(slog.LevelVar)
var logger *slog.Logger
var log *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))
	log = WithGroup("common")
}

func Logger() *slog.Logger {
	return logger
}

func WithGroup(groupName string) *slog.Logger {
	return logger.With(slog.Group("carbon", slog.String("pkg", groupName)))
}

func LogLevel(level slog.Level) {
	logLevel.Set(level)
}
