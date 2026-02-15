package entity

type Status struct {
	Level   uint32 `json:"level"`
	Defense uint32 `json:"defense"`
	Attack  uint32 `json:"attack"`

	Merchant  uint8 `json:"merchant"`
	Speed     uint8 `json:"speed"`
	Direction uint8 `json:"direction"`
	ChaosRate uint8 `json:"chaos_rate"`

	MaxHP uint32 `json:"max_hp"`
	MaxMP uint32 `json:"max_mp"`
	HP    uint32 `json:"hp"`
	MP    uint32 `json:"mp"`

	Str uint16 `json:"str"`
	Int uint16 `json:"int"`
	Dex uint16 `json:"dex"`
	Con uint16 `json:"con"`

	WMaster uint16 `json:"w_master"`
	FMaster uint16 `json:"f_master"`
	SMaster uint16 `json:"s_master"`
	TMaster uint16 `json:"t_master"`
}
