package main

import (
	"log"
	"os"

	"github.com/pessoal/wyd/apps/timer-server/client"
	"github.com/pessoal/wyd/packages/common/config"
)

func main() {
	cfg := config.Load()

	targetHost := "localhost"
	if h := os.Getenv("CONNECT_SERVER_HOST"); h != "" {
		targetHost = h
	}

	target := targetHost + ":" + cfg.GRPCPort
	log.Printf("Timer Server starting, connecting to %s", target)
	client.Start(target)
}
