package outgoing

import (
	"unsafe"

	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
)

type Message struct {
	Header  protocol.PacketHeader
	Message [96]byte
}

func NewMessage(message string) Message {
	var msgBytes [96]byte

	copy(msgBytes[:], message)
	return Message{
		Header: protocol.NewPacketHeader(
			protocol.PacketHeaderOptions{
				PacketID: protocol.PackageIDMessage,
				Size:     uint16(unsafe.Sizeof(Message{})),
				ClientID: 0,
			}),
		Message: msgBytes,
	}
}
