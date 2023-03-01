package entity

import (
	"go-oo/example16/bo"
)

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int

	Shop *Shop
}

func (ent *Item) Mapping() *bo.Item {
	boItem := new(bo.Item)
	boItem.ID = ent.ID
	boItem.Category = ent.Category
	boItem.Title = ent.Title
	boItem.Stock = ent.Stock
	boItem.PriceMarket = ent.PriceMarket
	boItem.Shop = ent.Shop.Mapping()

	boItem.Of()
	return boItem
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
