package bootstrap

import (
	"github.com/antonioducs/wyd/timer-server/internal/application"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc"
)

func NewUseCases(infra *InfraContainer, repos *RepositoryContainer) *application.UseCaseContainer {
	grpcPresenter := grpc.NewGRPCPresenter(infra.GRPCClient)

	return &application.UseCaseContainer{
		Login: usecase.NewLoginUsecase(grpcPresenter, repos.AccountReader),
	}
}
