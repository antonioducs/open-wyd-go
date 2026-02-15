package grpc

import (
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol/outgoing"
)

type GRPCPresenter struct {
	client *Client
}

func NewGRPCPresenter(c *Client) *GRPCPresenter {
	return &GRPCPresenter{client: c}
}

func (p *GRPCPresenter) SendMessage(sessionID uint32, text string) {
	pkt := outgoing.NewMessage(text)

	if p.client != nil {
		p.client.Send(sessionID, pkt.Header.PrepareToSend())
	}
}
