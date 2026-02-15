package bootstrap

import (
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/game"
)

type AppContainer struct {
	Router *game.Router
	Infra  *InfraContainer
}

func NewAppContainer() *AppContainer {

	infra := NewInfra()

	repos := NewRepositories(infra)

	usecases := NewUseCases(infra, repos)

	router := game.NewRouter(infra.Logger, usecases)

	infra.GRPCClient.SetHandler(router.RoutePacket)

	return &AppContainer{
		Router: router,
		Infra:  infra,
	}
}
