package protocol

import (
	"math/rand"
	"time"
	"unsafe"
)

type PacketHeader struct {
	Size      uint16
	Key       byte
	Checksum  byte
	PacketID  uint16
	ClientID  uint16
	TimeStamp uint32
}

type PacketHeaderOptions struct {
	PacketID uint16
	Size     uint16
	ClientID uint16
}

func NewPacketHeader(options PacketHeaderOptions) PacketHeader {
	return PacketHeader{
		Size:      options.Size,
		Key:       byte(rand.Intn(256)),
		Checksum:  0,
		PacketID:  options.PacketID,
		ClientID:  options.ClientID,
		TimeStamp: uint32(time.Now().Unix()),
	}
}

func (h *PacketHeader) PrepareToSend() []byte {
	ptr := unsafe.Pointer(h)

	totalSize := h.Size

	data := unsafe.Slice((*byte)(ptr), int(totalSize))

	return data
}
