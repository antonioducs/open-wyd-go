package routes

import (
	"net/http"

	"github.com/antonioducs/wyd/web-api/internal/api/handlers"
)

func AccountsRoutes(mux *http.ServeMux, h *handlers.AccountHandler) {
	mux.HandleFunc("POST /accounts", h.HandleCreate)
}
