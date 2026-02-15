package entity

import (
	"errors"
	"slices"
	"time"
)

type Character struct {
	ID        int32
	AccountID int32
	GuildID   int32
	Name      string
	Gold      int32
	Exp       uint64
	PosX      uint16
	PosY      uint16
	Slot      uint8
	ClassID   uint8

	Status    Status
	Equipment [16]Item
	Inventory [64]Item
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCharacter(accountID int32, name string, classID uint8, slotID uint8) (*Character, error) {
	if classID > 3 {
		return nil, errors.New("invalid class ID")
	}

	if slotID > 2 {
		return nil, errors.New("invalid slot ID")
	}

	character := &Character{
		AccountID: accountID,
		Name:      name,
		ClassID:   classID,
		PosX:      2105,
		PosY:      2038,
		Slot:      slotID,
		Status: Status{
			Level:   0,
			Defense: 100,
			Attack:  100,
			Speed:   6,
			MaxHP:   100,
			MaxMP:   100,
			HP:      100,
			MP:      100,
		},
	}

	switch classID {
	case 0:
		character.Status.Str = 8
		character.Status.Int = 4
		character.Status.Dex = 7
		character.Status.Con = 6
		character.Equipment[1] = Item{Index: 500, Effect1: 43, Effect1Value: 0}
		character.Equipment[2] = Item{Index: 602, Effect1: 43, Effect1Value: 0}
		character.Equipment[3] = Item{Index: 714, Effect1: 43, Effect1Value: 0}
		character.Equipment[4] = Item{Index: 826, Effect1: 43, Effect1Value: 0}
		character.Equipment[5] = Item{Index: 938, Effect1: 43, Effect1Value: 0}
	case 1:
		character.Status.Str = 5
		character.Status.Int = 8
		character.Status.Dex = 5
		character.Status.Con = 5
		character.Equipment[1] = Item{Index: 522, Effect1: 43, Effect1Value: 0}
		character.Equipment[2] = Item{Index: 630, Effect1: 43, Effect1Value: 0}
		character.Equipment[3] = Item{Index: 742, Effect1: 43, Effect1Value: 0}
		character.Equipment[4] = Item{Index: 854, Effect1: 43, Effect1Value: 0}
		character.Equipment[5] = Item{Index: 966, Effect1: 43, Effect1Value: 0}
	case 2:
		character.Status.Str = 6
		character.Status.Int = 6
		character.Status.Dex = 9
		character.Status.Con = 5
		character.Equipment[1] = Item{Index: 544, Effect1: 43, Effect1Value: 0}
		character.Equipment[2] = Item{Index: 658, Effect1: 43, Effect1Value: 0}
		character.Equipment[3] = Item{Index: 770, Effect1: 43, Effect1Value: 0}
		character.Equipment[4] = Item{Index: 882, Effect1: 43, Effect1Value: 0}
		character.Equipment[5] = Item{Index: 994, Effect1: 43, Effect1Value: 0}
	case 3:
		character.Status.Str = 8
		character.Status.Int = 9
		character.Status.Dex = 13
		character.Status.Con = 6
		character.Equipment[1] = Item{Index: 566, Effect1: 43, Effect1Value: 0}
		character.Equipment[2] = Item{Index: 686, Effect1: 43, Effect1Value: 0}
		character.Equipment[3] = Item{Index: 798, Effect1: 43, Effect1Value: 0}
		character.Equipment[4] = Item{Index: 910, Effect1: 43, Effect1Value: 0}
		character.Equipment[5] = Item{Index: 1022, Effect1: 43, Effect1Value: 0}
	}

	character.Equipment[0] = Item{Index: uint16(classID*10 + 1)}

	character.Inventory[0] = Item{Index: 401, Effect1: 61, Effect1Value: 40}
	character.Inventory[1] = Item{Index: 406, Effect1: 61, Effect1Value: 40}
	character.Inventory[60] = Item{Index: 3467}
	character.Inventory[61] = Item{Index: 3467}

	return character, nil
}

type CharacterList []*Character

func (list CharacterList) FindBySlot(slot uint8) *Character {
	index := slices.IndexFunc(list, func(c *Character) bool {
		return c.Slot == slot
	})
	if index == -1 {
		return nil
	}
	return list[index]
}
