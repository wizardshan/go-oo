package repository

import (
	"go-oo/example1/domain"
	"go-oo/example1/repository/entity"
)

type Item struct {
}

func NewItem() *Item {
	return new(Item)
}

func (repo *Item) Get() *domain.Item {

	item := new(entity.Item)
	item.ID = 1
	item.Category = 1
	item.Title = "T shirt1"
	item.Stock = 1
	item.PriceMarket = 100

	return item.Mapping()
}

func (repo *Item) All() domain.Items {

	var items entity.Items

	item1 := new(entity.Item)
	item1.ID = 1
	item1.Category = 1
	item1.Title = "T shirt1"
	item1.Stock = 1
	item1.PriceMarket = 100
	items = append(items, item1)

	item2 := new(entity.Item)
	item2.ID = 2
	item2.Category = 2
	item2.Title = "T shirt2"
	item2.Stock = 2
	item2.PriceMarket = 80
	items = append(items, item2)

	return items.Mapping()
}