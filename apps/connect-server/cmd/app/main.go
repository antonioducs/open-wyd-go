package main

import (
	"fmt"

	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/grpc"
	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/tcp"
	"github.com/antonioducs/wyd/logger"
	"github.com/antonioducs/wyd/pkg/configs"
)

func main() {
	cfg, err := configs.NewConfig()
	log := logger.NewLogger(cfg.Env)
	if err != nil {
		log.Error("Failed to load config", "error", err)
	}

	go grpc.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.GRPCPort))

	tcpServer := tcp.NewTCPServer(tcp.TCPServerOptions{
		Host:    cfg.Host,
		Port:    cfg.TCPPort,
		MaxConn: cfg.MaxConn,
		Logger:  log,
	})

	if err := tcpServer.Start(); err != nil {
		log.Error("Failed to start TCP server", "error", err)
	}
}
