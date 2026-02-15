package gateway

import (
	"context"

	"github.com/antonioducs/wyd/pkg/domain/entity"
)

type AccountReader interface {
	FindByUsername(ctx context.Context, username string) (*entity.Account, error)
}

type AccountWriter interface {
	Create(ctx context.Context, account *entity.Account) error
}

type AccountRepository interface {
	AccountReader
	AccountWriter
}
