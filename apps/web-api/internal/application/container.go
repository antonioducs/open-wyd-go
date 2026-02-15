package application

import "github.com/antonioducs/wyd/web-api/internal/domain/usecase"

type UseCaseContainer struct {
	CreateAccount *usecase.CreateAccountUsecase
}
