package bootstrap

import (
	"net/http"

	"github.com/antonioducs/wyd/pkg/configs"
	"github.com/antonioducs/wyd/web-api/internal/api"
	"github.com/antonioducs/wyd/web-api/internal/api/handlers"
	"github.com/antonioducs/wyd/web-api/internal/application"
)

func NewHTTPHandler(usecases *application.UseCaseContainer, config *configs.Config) http.Handler {
	accountHandler := handlers.NewAccountHandler(
		usecases,
	)

	appHandlers := &api.Handler{
		Accounts: accountHandler,
	}

	return api.NewRouter(appHandlers, config)
}
