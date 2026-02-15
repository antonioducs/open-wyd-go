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

func (r *PostgresAccountRepo) FindByUsername(ctx context.Context, username string) (*entity.Account, error) {
	dao, err := r.queries.GetAccountByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	acc := &entity.Account{
		ID:           uint32(dao.ID),
		Username:     dao.Username,
		PasswordHash: dao.PasswordHash,
		CreatedAt:    dao.CreatedAt.Time,
		Email:        dao.Email,
	}

	return acc, nil
}

func (r *PostgresAccountRepo) Create(ctx context.Context, account *entity.Account) error {
	model, err := r.queries.CreateAccount(ctx, db.CreateAccountParams{
		Username:     account.Username,
		PasswordHash: account.PasswordHash,
		Email:        account.Email,
	})
	if err != nil {
		return err
	}

	account.ID = uint32(model.ID)
	account.CreatedAt = model.CreatedAt.Time

	return nil
}
