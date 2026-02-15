package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/antonioducs/wyd/pkg/configs"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func NewLogger(env configs.Env) *slog.Logger {
	var handler slog.Handler

	if env == configs.Prod {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.StampMilli,
			NoColor:    !isatty.IsTerminal(os.Stdout.Fd()),
		})
	}

	return slog.New(handler)
}
