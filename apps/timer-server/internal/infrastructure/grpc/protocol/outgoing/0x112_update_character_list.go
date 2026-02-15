package outgoing

import (
	"unsafe"

	"github.com/antonioducs/wyd/pkg/domain/entity"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol"
	"github.com/antonioducs/wyd/timer-server/internal/infrastructure/grpc/protocol/models"
)

type UpdateCharacterList struct {
	Header     protocol.PacketHeader
	unk1       models.DWORD
	PosX       [4]models.WORD
	PosY       [4]models.WORD
	Name       [4][16]models.BYTE
	Status     [4]models.Status
	Equipment  [4][16]models.Item
	GuildIndex [4]models.WORD
	Gold       [4]models.DWORD
	Experience [4]models.DWORD
}

func nameToBytes(character *entity.Character) [16]byte {
	if character == nil {
		return [16]byte{}
	}
	return *(*[16]byte)(protocol.CopyTextToByte(16, character.Name))
}

func itemsToModel(character *entity.Character) [16]models.Item {
	if character == nil {
		return [16]models.Item{}
	}
	itemsModel := [16]models.Item{}
	for i, item := range character.Equipment {
		itemsModel[i] = models.NewItem(item)
	}
	return itemsModel
}

func posXToModel(character *entity.Character) models.WORD {
	if character == nil {
		return 0
	}
	return models.WORD(character.PosX)
}

func goldToModel(character *entity.Character) models.DWORD {
	if character == nil {
		return 0
	}
	return models.DWORD(character.Gold)
}

func experienceToModel(character *entity.Character) models.DWORD {
	if character == nil {
		return 0
	}
	return models.DWORD(character.Exp)
}

func posYToModel(character *entity.Character) models.WORD {
	if character == nil {
		return 0
	}
	return models.WORD(character.PosY)
}

func NewUpdateCharacterList(
	characters entity.CharacterList,
) UpdateCharacterList {
	character0 := characters.FindBySlot(0)
	character1 := characters.FindBySlot(1)
	character2 := characters.FindBySlot(2)
	character3 := characters.FindBySlot(3)

	return UpdateCharacterList{
		Header: protocol.NewPacketHeader(
			protocol.PacketHeaderOptions{
				PacketID: protocol.PackageUpdateCharacterList,
				Size:     uint16(unsafe.Sizeof(UpdateCharacterList{})),
				ClientID: 0,
			}),
		unk1: 0,
		PosX: [4]models.WORD{
			posXToModel(character0),
			posXToModel(character1),
			posXToModel(character2),
			posXToModel(character3),
		},
		PosY: [4]models.WORD{
			posYToModel(character0),
			posYToModel(character1),
			posYToModel(character2),
			posYToModel(character3),
		},
		Name: [4][16]models.BYTE{
			nameToBytes(character0),
			nameToBytes(character1),
			nameToBytes(character2),
			nameToBytes(character3),
		},
		Status: [4]models.Status{
			models.NewStatus(character0),
			models.NewStatus(character1),
			models.NewStatus(character2),
			models.NewStatus(character3),
		},
		Equipment: [4][16]models.Item{
			itemsToModel(character0),
			itemsToModel(character1),
			itemsToModel(character2),
			itemsToModel(character3),
		},
		GuildIndex: [4]models.WORD{0, 0, 0, 0},
		Gold: [4]models.DWORD{
			goldToModel(character0),
			goldToModel(character1),
			goldToModel(character2),
			goldToModel(character3),
		},
		Experience: [4]models.DWORD{
			experienceToModel(character0),
			experienceToModel(character1),
			experienceToModel(character2),
			experienceToModel(character3),
		},
	}
}
