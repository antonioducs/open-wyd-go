package api

import (
	"fmt"
	"net/http"

	"github.com/antonioducs/wyd/pkg/configs"
	_ "github.com/antonioducs/wyd/web-api/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/antonioducs/wyd/web-api/internal/api/handlers"
	"github.com/antonioducs/wyd/web-api/internal/api/routes"
)

type Handler struct {
	Accounts *handlers.AccountHandler
}

func NewRouter(h *Handler, config *configs.Config) *http.ServeMux {
	apiV1 := http.NewServeMux()
	mainMux := http.NewServeMux()

	routes.AccountsRoutes(apiV1, h.Accounts)

	mainMux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1))

	mainMux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	mainMux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://%s:%s/swagger/doc.json", config.HTTPHost, config.HTTPPort)),
	))

	return mainMux
}
