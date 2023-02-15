package entity

import (
	"go-oo/example2/bo"
)

type Items []*Item

type Item struct {
	ID          int
	Title       string
	Stock       int
	PriceMarket int
}

func (ent *Item) Mapping() *bo.Item {
	bo := new(bo.Item)
	bo.ID = ent.ID
	bo.Title = ent.Title
	bo.Stock = ent.Stock
	bo.PriceMarket = ent.PriceMarket
	bo.PriceCalculate()
	return bo

}

func (ents Items) Mapping() bo.Items {
	entsLen := len(ents)
	bos := make(bo.Items, entsLen)
	if entsLen > 0 {
		for entIndex := 0; entIndex < entsLen; entIndex++ {
			bos[entIndex] = ents[entIndex].Mapping()
		}
		return bos
	}

	return nil
}
