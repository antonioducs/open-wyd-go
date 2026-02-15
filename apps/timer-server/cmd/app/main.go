package main

import (
	"github.com/antonioducs/wyd/timer-server/cmd/app/bootstrap"
)

func main() {
	container := bootstrap.NewAppContainer()

	container.Infra.GRPCClient.Start()
}
