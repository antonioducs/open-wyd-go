package main

import (
	"github.com/antonioducs/wyd/connect-server/configs"
	"github.com/antonioducs/wyd/connect-server/internal/infrastructure/tcp"
	"github.com/antonioducs/wyd/logger"
)

func main() {
	cfg, err := configs.Load()
	log := logger.New(cfg.Env)
	if err != nil {
		log.Error("Failed to load config", "error", err)
	}

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
