package bootstrap

import (
	"github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/pkg/storage/postgres/repository"
)

type RepositoryContainer struct {
	AccountReader gateway.AccountReader
}

func NewRepositories(infra *InfraContainer) *RepositoryContainer {
	return &RepositoryContainer{
		AccountReader: repository.NewPostgresAccountRepo(infra.DB),
	}
}
