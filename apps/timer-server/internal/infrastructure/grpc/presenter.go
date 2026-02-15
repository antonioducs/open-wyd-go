package grpc

import (
	"github.com/antonioducs/wyd/pkg/domain/entity"
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

func (p *GRPCPresenter) SendCharList(sessionID uint32) {
	pkt := outgoing.NewCharList()

	if p.client != nil {
		p.client.Send(sessionID, pkt.Header.PrepareToSend())
	}
}

func (p *GRPCPresenter) SendUpdateCharacterList(sessionID uint32, characters entity.CharacterList) {
	pkt := outgoing.NewUpdateCharacterList(characters)

	if p.client != nil {
		p.client.Send(sessionID, pkt.Header.PrepareToSend())
	}
}
