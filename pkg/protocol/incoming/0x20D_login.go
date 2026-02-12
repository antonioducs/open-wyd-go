package incoming

import (
	"bytes"

	"github.com/antonioducs/wyd/protocol"
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
