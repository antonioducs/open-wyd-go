package main

import (
	"sync"
	"time"

	"github.com/pessoal/wyd/apps/connect-server/grpc"
	"github.com/pessoal/wyd/apps/connect-server/tcp"
	"github.com/pessoal/wyd/packages/common/config"
)

func main() {
	cfg := config.Load()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		tcp.Start(cfg.Port)
	}()

	go func() {
		defer wg.Done()
		// Wait a bit to ensure logs don't mix up too much
		time.Sleep(100 * time.Millisecond)
		grpc.Start(cfg.GRPCPort)
	}()

	wg.Wait()
}
