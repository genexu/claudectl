package utils

import (
	"context"
	"log/slog"
	"os"

	"go.uber.org/fx"
)

type Logger struct {
	*slog.Logger
	logFile *os.File
}

func getLogLevel(debugFlag bool) slog.Level {
	if debugFlag {
		return slog.LevelDebug
	}
	return slog.LevelInfo
}

func NewLogger(lc fx.Lifecycle, debugFlag bool) (*Logger, error) {
	logFile, err := os.OpenFile("claudectl.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: getLogLevel(debugFlag),
	})

	logger := &Logger{
		Logger:  slog.New(handler),
		logFile: logFile,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("logger initialized")
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("closing log file")
			return logFile.Close()
		},
	})

	return logger, nil
}
