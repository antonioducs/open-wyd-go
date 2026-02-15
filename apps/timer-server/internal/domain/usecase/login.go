package usecase

import (
	"context"

	shared_gateway "github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	output        gateway.GameOutput
	accountReader shared_gateway.AccountReader
}

func NewLoginUsecase(
	output gateway.GameOutput,
	accountReader shared_gateway.AccountReader,
) *LoginUsecase {
	return &LoginUsecase{
		output:        output,
		accountReader: accountReader,
	}
}

type LoginInput struct {
	Context   context.Context
	SessionID uint32
	Username  string
	Password  string
}

func (u *LoginUsecase) Execute(input LoginInput) {

	account, err := u.accountReader.FindByUsername(input.Context, input.Username)
	if err != nil {
		u.output.SendMessage(input.SessionID, "Usuario e/ou senha incorretos")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.PasswordHash), []byte(input.Password))
	if err != nil {
		u.output.SendMessage(input.SessionID, "Usuario e/ou senha incorretos")
		return
	}

	u.output.SendCharList(input.SessionID)
}
