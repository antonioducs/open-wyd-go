package usecase

import (
	"context"
	"errors"
	"log/slog"

	"github.com/antonioducs/wyd/pkg/domain/entity"
	"github.com/antonioducs/wyd/pkg/domain/gateway"
	"golang.org/x/crypto/bcrypt"
)

type CreateAccountUsecase struct {
	accountRepository gateway.AccountRepository
	logger            *slog.Logger
}

type CreateAccountOutput struct {
	ID       uint32
	Username string
}

type CreateAccountInput struct {
	Username string
	Password string
	Email    string
}

func NewCreateAccountUsecase(
	accountRepository gateway.AccountRepository,
	logger *slog.Logger,
) *CreateAccountUsecase {
	return &CreateAccountUsecase{
		accountRepository: accountRepository,
		logger:            logger,
	}
}

func (uc *CreateAccountUsecase) Execute(
	ctx context.Context,
	input CreateAccountInput,
) (*CreateAccountOutput, error) {
	exists, _ := uc.accountRepository.FindByUsername(ctx, input.Username)

	if exists != nil {
		return nil, errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.logger.Error("failed to hash password", "error", err)
		return nil, errors.New("failed to hash password")
	}

	account := &entity.Account{
		Username:     input.Username,
		Email:        input.Email,
		PasswordHash: string(hashedPassword),
	}

	err = uc.accountRepository.Create(ctx, account)
	if err != nil {
		uc.logger.Error("failed to create account", "error", err)
		return nil, errors.New("failed to create account")
	}

	return &CreateAccountOutput{
		ID:       account.ID,
		Username: account.Username,
	}, nil
}
