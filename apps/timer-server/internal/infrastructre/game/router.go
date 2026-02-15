package game

import (
	"encoding/binary"
	"fmt"
	"log/slog"

	"github.com/antonioducs/wyd/timer-server/internal/application"
	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc/protocol"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc/protocol/incoming"
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

func (r *Router) RoutePacket(sessionID uint32, payload []byte) {
	packetID := binary.LittleEndian.Uint16(payload[4:6])

	r.logger.Debug("Pacote Recebido", "id", fmt.Sprintf("0x%X", packetID))

	switch packetID {
	case protocol.PackageIDLogin:
		msg := protocol.Cast[incoming.Login](payload)
		r.usecases.Login.Execute(usecase.LoginInput{
			SessionID: sessionID,
			Username:  msg.GetUsername(),
			Password:  msg.GetPassword(),
		})

	default:
		r.logger.Warn("Pacote desconhecido", "id", packetID)
	}
}
