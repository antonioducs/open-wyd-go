package bootstrap

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/antonioducs/wyd/logger"
	"github.com/antonioducs/wyd/pkg/configs"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfraContainer struct {
	Config     *configs.Config
	Logger     *slog.Logger
	DB         *pgxpool.Pool
	GRPCClient *grpc.Client
}

func NewInfra() *InfraContainer {
	config, err := configs.NewConfig()
	if err != nil {
		panic("falha fatal ao carregar config: " + err.Error())
	}
	logger := logger.NewLogger(config.Env)

	db, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		panic("falha fatal ao conectar no banco: " + err.Error())
	}

	grpcClient := grpc.NewClient(
		fmt.Sprintf("%s:%s", config.Host, config.GRPCPort),
		logger,
		nil)

	return &InfraContainer{
		Config:     config,
		Logger:     logger,
		DB:         db,
		GRPCClient: grpcClient,
	}
}
