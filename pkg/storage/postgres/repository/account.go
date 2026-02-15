package repository

import (
	"context"

	"github.com/antonioducs/wyd/pkg/domain/entity"

	"github.com/antonioducs/wyd/pkg/storage/postgres/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresAccountRepo struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func NewPostgresAccountRepo(pool *pgxpool.Pool) *PostgresAccountRepo {
	return &PostgresAccountRepo{
		queries: db.New(pool),
		pool:    pool,
	}
}

func (r *PostgresAccountRepo) FindByUsername(username string) (*entity.Account, error) {
	dao, err := r.queries.GetAccountByUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}

	acc := &entity.Account{
		ID:           uint32(dao.ID),
		Username:     dao.Username,
		PasswordHash: dao.PasswordHash,
	}

	return acc, nil
}
