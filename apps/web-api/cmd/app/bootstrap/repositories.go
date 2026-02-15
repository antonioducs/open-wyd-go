package bootstrap

import (
	"github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/pkg/storage/postgres/repository"
)

type RepositoryContainer struct {
	Account gateway.AccountRepository
}

func NewRepositories(infra *InfraContainer) *RepositoryContainer {
	return &RepositoryContainer{
		Account: repository.NewPostgresAccountRepo(infra.DB),
	}
}
