package bootstrap

import (
	shared_gateway "github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/pkg/storage/postgres/repository"
	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/game/session"
)

type RepositoryContainer struct {
	AccountReader       shared_gateway.AccountReader
	CharacterRepository shared_gateway.CharacterRepository
	SessionRepository   gateway.SessionRepository
}

func NewRepositories(infra *InfraContainer) *RepositoryContainer {
	return &RepositoryContainer{
		AccountReader:       repository.NewPostgresAccountRepo(infra.DB),
		CharacterRepository: repository.NewPostgresCharacterRepo(infra.DB),
		SessionRepository:   session.NewManager(),
	}
}
