package gateway

import "github.com/antonioducs/wyd/pkg/domain/entity"

type GameOutput interface {
	SendMessage(sessionID uint32, message string)
	SendCharList(sessionID uint32)
	SendUpdateCharacterList(sessionID uint32, characters entity.CharacterList)
}
