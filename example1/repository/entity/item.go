package entity

import (
	"go-oo/example1/domain"
)

type Items []*Item

type Item struct {
	ID       int
	Category int
	Title    string
	Stock    int

	PriceMarket int
}

func (ent *Item) Mapping() *domain.Item {
	/**************** mapping start ****************/
	dom := new(domain.Item)
	dom.ID = ent.ID
	dom.Category = ent.Category
	dom.Title = ent.Title
	dom.Stock = ent.Stock
	dom.PriceMarket = ent.PriceMarket
	return dom

	/**************** mapping end  ****************/
}

func (ent Items) Mapping() domain.Items {
	/**************** mapping start ****************/
	entItemsLen := len(ent)
	dom := make(domain.Items, entItemsLen)
	if entItemsLen > 0 {
		for entItemsIndex := 0; entItemsIndex < entItemsLen; entItemsIndex++ {
			dom[entItemsIndex] = ent[entItemsIndex].Mapping()
		}
		return dom
	}

	return nil
	/**************** mapping end  ****************/
}
