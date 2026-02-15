package repository

import (
	"context"
	"encoding/json"

	"github.com/antonioducs/wyd/pkg/domain/entity"
	"github.com/antonioducs/wyd/pkg/storage/postgres/db"
	"github.com/antonioducs/wyd/pkg/storage/postgres/mappers"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCharacterRepo struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func NewPostgresCharacterRepo(pool *pgxpool.Pool) *PostgresCharacterRepo {
	return &PostgresCharacterRepo{
		queries: db.New(pool),
		pool:    pool,
	}
}

func (r *PostgresCharacterRepo) Create(ctx context.Context, character *entity.Character) error {

	statusBin, _ := json.Marshal(character.Status)
	equipmentBin, _ := json.Marshal(character.Equipment)
	inventoryBin, _ := json.Marshal(character.Inventory)

	model, err := r.queries.CreateCharacter(ctx, db.CreateCharacterParams{
		AccountID:  character.AccountID,
		Name:       character.Name,
		Level:      int32(character.Status.Level),
		Experience: int64(character.Exp),
		Gold:       character.Gold,
		GuildID:    int32(character.GuildID),
		PosX:       int32(character.PosX),
		PosY:       int32(character.PosY),
		Status:     statusBin,
		Equipment:  equipmentBin,
		Inventory:  inventoryBin,
	})
	if err != nil {
		return err
	}

	character.ID = model.ID
	character.CreatedAt = model.CreatedAt.Time
	character.UpdatedAt = model.UpdatedAt.Time

	return nil
}

func (r *PostgresCharacterRepo) FindByID(ctx context.Context, id int32) (*entity.Character, error) {
	model, err := r.queries.GetCharacterByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mappers.ToDomain(model)
}

func (r *PostgresCharacterRepo) FindByAccountID(ctx context.Context, accountID int32) ([]*entity.Character, error) {
	models, err := r.queries.GetCharactersByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return []*entity.Character{}, nil
	}

	characters := make([]*entity.Character, len(models))
	for i, model := range models {
		character, err := mappers.ToDomain(model)
		if err != nil {
			return nil, err
		}
		characters[i] = character
	}

	return characters, nil
}

func (r *PostgresCharacterRepo) Update(ctx context.Context, character *entity.Character) error {
	statusBin, _ := json.Marshal(character.Status)
	equipmentBin, _ := json.Marshal(character.Equipment)
	inventoryBin, _ := json.Marshal(character.Inventory)

	return r.queries.UpdateCharacter(ctx, db.UpdateCharacterParams{
		ID:         character.ID,
		Name:       character.Name,
		Level:      int32(character.Status.Level),
		Experience: int64(character.Exp),
		Gold:       character.Gold,
		GuildID:    int32(character.GuildID),
		PosX:       int32(character.PosX),
		PosY:       int32(character.PosY),
		Status:     statusBin,
		Equipment:  equipmentBin,
		Inventory:  inventoryBin,
	})
}

func (r *PostgresCharacterRepo) Delete(ctx context.Context, id int32) error {
	return r.queries.DeleteCharacter(ctx, id)
}

func (r *PostgresCharacterRepo) FindByName(ctx context.Context, name string) (*entity.Character, error) {
	model, err := r.queries.GetCharacterByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return mappers.ToDomain(model)
}
