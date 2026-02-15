package bootstrap

import (
	"net/http"
)

type AppContainer struct {
	Handler http.Handler
	Infra   *InfraContainer
}

func NewAppContainer() *AppContainer {
	infra := NewInfra()

	repos := NewRepositories(infra)

	usecases := NewUseCases(repos, infra.Logger)

	httpHandler := NewHTTPHandler(usecases, infra.Config)

	return &AppContainer{
		Handler: httpHandler,
		Infra:   infra,
	}
}
