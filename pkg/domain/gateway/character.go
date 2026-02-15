package gateway

import (
	"context"

	"github.com/antonioducs/wyd/pkg/domain/entity"
)

type CharacterRepository interface {
	Create(ctx context.Context, character *entity.Character) error
	Update(ctx context.Context, character *entity.Character) error
	Delete(ctx context.Context, id int32) error
	FindByAccountID(ctx context.Context, accountID int32) ([]*entity.Character, error)
	FindByID(ctx context.Context, id int32) (*entity.Character, error)
	FindByName(ctx context.Context, name string) (*entity.Character, error)
}
