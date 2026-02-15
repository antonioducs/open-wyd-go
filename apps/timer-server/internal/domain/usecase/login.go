package usecase

import (
	shared_gateway "github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
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
	SessionID uint32
	Username  string
	Password  string
}

func (u *LoginUsecase) Execute(input LoginInput) {

	account, err := u.accountReader.FindByUsername(input.Username)
	if err != nil {
		u.output.SendMessage(input.SessionID, "Usuario e/ou senha incorretos")
		return
	}

	if account.PasswordHash != input.Password {
		u.output.SendMessage(input.SessionID, "Usuario e/ou senha incorretos")
		return
	}

	u.output.SendMessage(input.SessionID, "Login realizado com sucesso")
}
