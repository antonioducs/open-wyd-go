package mappers

import (
	"encoding/json"

	"github.com/antonioducs/wyd/pkg/domain/entity"
	"github.com/antonioducs/wyd/pkg/storage/postgres/db"
)

func ToDomain(row db.Character) (*entity.Character, error) {

	c := &entity.Character{
		ID:        row.ID,
		AccountID: row.AccountID,
		Name:      row.Name,
		Exp:       uint64(row.Experience),
		Gold:      row.Gold,
		GuildID:   row.GuildID,
		PosX:      uint16(row.PosX),
		PosY:      uint16(row.PosY),
		Slot:      uint8(row.Slot),
		ClassID:   uint8(row.ClassID),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}

	if len(row.Status) > 0 {
		if err := json.Unmarshal(row.Status, &c.Status); err != nil {
			return nil, err
		}
	}

	if len(row.Equipment) > 0 {
		if err := json.Unmarshal(row.Equipment, &c.Equipment); err != nil {
			return nil, err
		}
	}

	if len(row.Inventory) > 0 {
		if err := json.Unmarshal(row.Inventory, &c.Inventory); err != nil {
			return nil, err
		}
	}

	return c, nil
}
