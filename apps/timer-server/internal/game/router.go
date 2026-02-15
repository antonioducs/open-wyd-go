package game

import (
	"encoding/binary"
	"fmt"
	"log/slog"

	"github.com/antonioducs/wyd/timer-server/internal/domain/usecase"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc/protocol"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructre/grpc/protocol/incoming"
)

type Router struct {
	logger    *slog.Logger
	client    *grpc.Client
	presenter *grpc.GRPCPresenter
}

func NewRouter(logger *slog.Logger) *Router {
	return &Router{
		logger: logger,
	}
}

func (r *Router) SetClient(c *grpc.Client) {
	r.client = c
}

func (r *Router) SetPresenter(p *grpc.GRPCPresenter) {
	r.presenter = p
}

func (r *Router) RoutePacket(sessionID uint32, payload []byte) {
	packetID := binary.LittleEndian.Uint16(payload[4:6])

	r.logger.Debug("Pacote Recebido", "id", fmt.Sprintf("0x%X", packetID))

	switch packetID {
	case protocol.PackageIDLogin:
		msg := protocol.Cast[incoming.Login](payload)
		usecase.NewLoginUsecase(r.presenter).Execute(usecase.LoginInput{
			SessionID: sessionID,
			Username:  msg.GetUsername(),
			Password:  msg.GetPassword(),
		})

	default:
		r.logger.Warn("Pacote desconhecido", "id", packetID)
	}
}
