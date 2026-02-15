package main

import (
	"fmt"

	"github.com/antonioducs/wyd/logger"
	"github.com/antonioducs/wyd/pkg/configs"
	"github.com/antonioducs/wyd/timer-server/internal/game"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc"
)

func main() {
	config, err := configs.NewConfig()
	log := logger.NewLogger(config.Env)
	if err != nil {
		log.Error("Failed to load config", "error", err)
	}

	router := game.NewRouter(log)

	client := grpc.NewClient(
		fmt.Sprintf("%s:%s", config.Host, config.GRPCPort),
		log,
		router.RoutePacket,
	)

	presenter := grpc.NewGRPCPresenter(client)

	router.SetPresenter(presenter)
	router.SetClient(client)

	client.Start()
}
