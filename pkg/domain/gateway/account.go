package gateway

import "github.com/antonioducs/wyd/pkg/domain/entity"

type AccountReader interface {
	FindByUsername(username string) (*entity.Account, error)
}
