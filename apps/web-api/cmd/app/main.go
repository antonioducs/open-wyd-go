package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antonioducs/wyd/web-api/cmd/app/bootstrap"
)

// @title           Open WYD Web API
// @version         1.0
// @description     API for managing accounts and characters of the Open WYD server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	app := bootstrap.NewAppContainer()
	logger := app.Infra.Logger

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.Infra.Config.HTTPPort),
		Handler:      app.Handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Info("🚀 Web API rodando na porta 8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Falha fatal no servidor HTTP", "err", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Erro ao desligar servidor forçado", "err", err)
	}

	app.Infra.DB.Close()
	logger.Info("Servidor desligado com sucesso. Tchau! 👋")
}
