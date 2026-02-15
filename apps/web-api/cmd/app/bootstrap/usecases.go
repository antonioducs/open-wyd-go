package bootstrap

import (
	"log/slog"

	"github.com/antonioducs/wyd/web-api/internal/application"
	"github.com/antonioducs/wyd/web-api/internal/domain/usecase"
)

func NewUseCases(repos *RepositoryContainer, logger *slog.Logger) *application.UseCaseContainer {
	return &application.UseCaseContainer{
		CreateAccount: usecase.NewCreateAccountUsecase(repos.Account, logger),
	}
}
