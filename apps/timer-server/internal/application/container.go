package application

import (
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/character"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/login"
)

type UseCaseContainer struct {
	Login           *login.LoginUsecase
	CreateCharacter *character.CreateCharacterUseCase
}
