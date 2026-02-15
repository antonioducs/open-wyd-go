package application

import "github.com/antonioducs/wyd/timer-server/internal/domain/usecase"

type UseCaseContainer struct {
	Login *usecase.LoginUsecase
}
