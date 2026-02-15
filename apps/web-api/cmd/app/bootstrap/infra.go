package bootstrap

import (
	"context"
	"log/slog"

	"github.com/antonioducs/wyd/logger"
	"github.com/antonioducs/wyd/pkg/configs"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfraContainer struct {
	Config *configs.Config
	Logger *slog.Logger
	DB     *pgxpool.Pool
}

func NewInfra() *InfraContainer {
	config, err := configs.NewConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger := logger.NewLogger(config.Env)

	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return &InfraContainer{
		Config: config,
		Logger: logger,
		DB:     pool,
	}
}
