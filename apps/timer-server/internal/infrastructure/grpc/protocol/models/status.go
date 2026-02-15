package models

import "github.com/antonioducs/wyd/pkg/domain/entity"

type Status struct {
	Level           DWORD
	DefensePoints   DWORD
	AttackPoints    DWORD
	Merchant        BYTE
	Speed           BYTE
	Direction       BYTE
	ChaosRate       BYTE
	MaxHealthPoints DWORD
	MaxMagicPoints  DWORD
	HealthPoints    DWORD
	MagicPoints     DWORD
	Strength        WORD
	Intelligence    WORD
	Dexterity       WORD
	Constitution    WORD
	WMaster         WORD
	FMaster         WORD
	SMaster         WORD
	TMaster         WORD
}

func NewStatus(character *entity.Character) Status {
	if character == nil {
		return Status{}
	}
	status := character.Status
	return Status{
		Level:           status.Level,
		DefensePoints:   status.Defense,
		AttackPoints:    status.Attack,
		Merchant:        status.Merchant,
		Speed:           status.Speed,
		Direction:       status.Direction,
		ChaosRate:       status.ChaosRate,
		MaxHealthPoints: status.MaxHP,
		MaxMagicPoints:  status.MaxMP,
		HealthPoints:    status.HP,
		MagicPoints:     status.MP,
		Strength:        status.Str,
		Intelligence:    status.Int,
		Dexterity:       status.Dex,
		Constitution:    status.Con,
		WMaster:         status.WMaster,
		FMaster:         status.FMaster,
		SMaster:         status.SMaster,
		TMaster:         status.TMaster,
	}
}
