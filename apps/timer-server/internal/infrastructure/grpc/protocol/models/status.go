package models

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
