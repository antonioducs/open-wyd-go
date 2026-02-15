package entity

type Item struct {
	Index        uint16 `json:"index"`
	Effect1      uint8  `json:"effect1"`
	Effect1Value uint8  `json:"effect1_value"`
	Effect2      uint8  `json:"effect2"`
	Effect2Value uint8  `json:"effect2_value"`
	Effect3      uint8  `json:"effect3"`
	Effect3Value uint8  `json:"effect3_value"`
}

func (i Item) IsEmpty() bool {
	return i.Index == 0
}
