package incoming

import (
	"bytes"

	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
)

type CreateChar struct {
	Header  protocol.PacketHeader
	SlotID  uint32
	Name    [12]byte
	unk1    int32
	ClassID byte
}

func (c *CreateChar) GetName() string {
	return string(bytes.Trim(c.Name[:], "\x00"))
}

func NewCreateChar(data []byte) *CreateChar {
	return protocol.Cast[CreateChar](data)
}
