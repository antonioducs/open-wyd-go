package login

import (
	"context"

	shared_gateway "github.com/antonioducs/wyd/pkg/domain/gateway"
	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	output              gateway.GameOutput
	accountReader       shared_gateway.AccountReader
	sessionRepository   gateway.SessionRepository
	characterRepository shared_gateway.CharacterRepository
}

func NewLoginUsecase(
	output gateway.GameOutput,
	accountReader shared_gateway.AccountReader,
	characterRepository shared_gateway.CharacterRepository,
	sessionRepository gateway.SessionRepository,
) *LoginUsecase {
	return &LoginUsecase{
		output:              output,
		accountReader:       accountReader,
		characterRepository: characterRepository,
		sessionRepository:   sessionRepository,
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

	characters, err := u.characterRepository.FindByAccountID(input.Context, account.ID)
	if err != nil {
		u.output.SendMessage(input.SessionID, "Erro ao buscar personagens")
		return
	}

	u.sessionRepository.Add(input.SessionID, account, characters)

	u.output.SendCharList(input.SessionID)
}
