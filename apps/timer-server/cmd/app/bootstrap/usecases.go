package bootstrap

import (
	"github.com/antonioducs/wyd/timer-server/internal/application"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/character"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/login"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc"
)

func NewUseCases(infra *InfraContainer, repos *RepositoryContainer) *application.UseCaseContainer {
	grpcPresenter := grpc.NewGRPCPresenter(infra.GRPCClient)

	return &application.UseCaseContainer{
		Login: login.NewLoginUsecase(
			grpcPresenter,
			repos.AccountReader,
			repos.CharacterRepository,
			repos.SessionRepository,
		),
		CreateCharacter: character.NewCreateCharacterUseCase(
			grpcPresenter,
			repos.CharacterRepository,
			repos.SessionRepository,
			infra.Logger,
		),
	}
}
