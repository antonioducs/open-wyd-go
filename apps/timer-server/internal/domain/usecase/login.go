package usecase

import (
	"fmt"

	"github.com/antonioducs/wyd/timer-server/internal/domain/gateway"
)

type LoginUsecase struct {
	output gateway.GameOutput
}

func NewLoginUsecase(output gateway.GameOutput) *LoginUsecase {
	return &LoginUsecase{output: output}
}

type LoginInput struct {
	SessionID uint32
	Username  string
	Password  string
}

func (u *LoginUsecase) Execute(input LoginInput) {

	fmt.Println(input)

	u.output.SendMessage(input.SessionID, "Hello, "+input.Username)
}
