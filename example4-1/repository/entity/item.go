package entity

import (
	"go-oo/example4/domain"
)

type Items []*Item

type Item struct {
	ID    int
	Category int
	Title string
	Stock int

	PriceMarket int
}

func (ent *Item) Mapping() *domain.Item {
	dom := new(domain.Item)
	dom.ID = ent.ID
	dom.Category = ent.Category
	dom.Title = ent.Title
	dom.Stock = ent.Stock
	dom.PriceMarket = ent.PriceMarket
	return dom

}

func (ent Items) Mapping() domain.Items {
	entItemsLen := len(ent)
	dom := make(domain.Items, entItemsLen)
	if entItemsLen > 0 {
		for entItemsIndex := 0; entItemsIndex < entItemsLen; entItemsIndex++ {
			dom[entItemsIndex] = ent[entItemsIndex].Mapping()
		}
		return dom
	}

	return nil
}
