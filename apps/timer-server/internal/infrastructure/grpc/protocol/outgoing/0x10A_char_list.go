package outgoing

import (
	"unsafe"

	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol/models"
)

type CharList struct {
	Header     protocol.PacketHeader
	unk1       [20]models.BYTE
	PosX       [4]models.WORD
	PosY       [4]models.WORD
	Names      [4][16]models.BYTE
	Status     [4]models.Status
	Equipments [4][16]models.Item
	GuildIndex [4]models.WORD
	Gold       [4]models.DWORD
	Experience [4]models.DWORD
	Unused2    [4]models.DWORD
	Storage    [120]models.Item
}

func NewCharList() CharList {
	return CharList{
		Header: protocol.NewPacketHeader(
			protocol.PacketHeaderOptions{
				PacketID: protocol.PackageIDCharList,
				Size:     uint16(unsafe.Sizeof(CharList{})),
				ClientID: 0,
			}),
	}
}
