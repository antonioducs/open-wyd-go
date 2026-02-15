package incoming

import (
	"bytes"

	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
)

type Login struct {
	Header   protocol.PacketHeader
	Password [10]byte
	unk      uint16
	Username [12]byte
}

func (l *Login) GetUsername() string {
	return string(bytes.Trim(l.Username[:], "\x00"))
}

func (l *Login) GetPassword() string {
	return string(bytes.Trim(l.Password[:], "\x00"))
}

func NewLogin(data []byte) *Login {
	return protocol.Cast[Login](data)
}
