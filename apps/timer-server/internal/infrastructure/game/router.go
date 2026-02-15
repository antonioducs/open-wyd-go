package game

import (
	"context"
	"encoding/binary"
	"fmt"
	"log/slog"

	"github.com/antonioducs/wyd/timer-server/internal/application"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/character"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase/login"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol/incoming"
)

type Router struct {
	logger   *slog.Logger
	usecases *application.UseCaseContainer
}

func NewRouter(logger *slog.Logger, usecases *application.UseCaseContainer) *Router {
	return &Router{
		logger:   logger,
		usecases: usecases,
	}
}

func (r *Router) RoutePacket(ctx context.Context, sessionID uint32, payload []byte) {
	packetID := binary.LittleEndian.Uint16(payload[4:6])

	r.logger.Debug("Pacote Recebido", "id", fmt.Sprintf("0x%X", packetID))

	switch packetID {
	case protocol.PackageIDLogin:
		msg := protocol.Cast[incoming.Login](payload)

		r.usecases.Login.Execute(login.LoginInput{
			Context:   ctx,
			SessionID: sessionID,
			Username:  msg.GetUsername(),
			Password:  msg.GetPassword(),
		})

	case protocol.PackageIDCreateCharacter:
		msg := protocol.Cast[incoming.CreateChar](payload)
		r.usecases.CreateCharacter.Execute(character.CreateCharacterInput{
			Context:   ctx,
			SessionID: sessionID,
			Name:      msg.GetName(),
			ClassID:   msg.ClassID,
			SlotID:    uint8(msg.SlotID),
		})

	default:
		r.logger.Warn("Pacote desconhecido", "id", packetID)
	}
}
