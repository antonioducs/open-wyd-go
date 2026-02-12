package logger

import (
	"log/slog"
	"os"

	"github.com/antonioducs/wyd/connect-server/configs"
)

func New(env configs.Env) *slog.Logger {
	var handler slog.Handler

	if env == configs.Prod {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}

	return slog.New(handler)
}
