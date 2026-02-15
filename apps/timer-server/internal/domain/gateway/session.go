package gateway

import (
	"github.com/antonioducs/wyd/pkg/domain/entity"
	game "github.com/antonioducs/wyd/timer-server/internal/infrastructure/game/session"
)

type SessionRepository interface {
	Add(sessionID uint32, acc *entity.Account, chars []*entity.Character)
	Get(sessionID uint32) (*game.UserSession, bool)
	Remove(sessionID uint32)
}
