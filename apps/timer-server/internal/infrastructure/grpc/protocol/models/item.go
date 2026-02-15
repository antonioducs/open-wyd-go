package models

import "github.com/antonioducs/wyd/pkg/domain/entity"

type Item struct {
	Index        WORD
	Effect1      BYTE
	Effect1Value BYTE
	Effect2      BYTE
	Effect2Value BYTE
	Effect3      BYTE
	Effect3Value BYTE
}

func NewItem(item entity.Item) Item {
	return Item{
		Index:        item.Index,
		Effect1:      item.Effect1,
		Effect1Value: item.Effect1Value,
		Effect2:      item.Effect2,
		Effect2Value: item.Effect2Value,
		Effect3:      item.Effect3,
		Effect3Value: item.Effect3Value,
	}
}
