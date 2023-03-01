package repository

import (
	"go-oo/example13/bo"
	"go-oo/example13/repository/entity"
)

type Item struct {
}

func NewItem() *Item {
	return new(Item)
}

func (repo *Item) Get() *bo.Item {

	item := new(entity.Item)
	item.ID = 1
	item.Category = 1
	item.Title = "T shirt1"
	item.Stock = 1
	item.PriceMarket = 100

	return item.Mapping()
}

func (repo *Item) All() bo.Items {

	items := entity.Items{}

	item1 := new(entity.Item)
	item1.ID = 1
	item1.Category = 1
	item1.Title = "T shirt1"
	item1.Stock = 1
	item1.PriceMarket = 100
	items = append(items, item1)

	return items.Mapping()
}
