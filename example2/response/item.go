package response

import "go-oo/example2/domain"

type Items []*Item

type Item struct {
	ID          int    `json:"id"`
	Category    int    `json:"category"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	PriceMarket int    `json:"priceMarket"`
	Price       int    `json:"price"`
}

func (resp *Item) Mapping(dom *domain.Item) {
	/**************** mapping start ****************/
	resp.ID = dom.ID
	resp.Category = dom.Category
	resp.Title = dom.Title
	resp.Stock = dom.Stock
	resp.PriceMarket = dom.PriceMarket

	instance := dom.OfInstance()
	if priceCalculator, ok := instance.(domain.ItemPriceCalculator); ok {
		resp.Price = priceCalculator.Price()
	}

	/**************** mapping end  ****************/
}

func (resp *Items) Mapping(dom domain.Items) {
	/**************** mapping start ****************/
	domItemsLen := len(dom)
	*resp = make(Items, domItemsLen)
	if domItemsLen > 0 {
		for domItemsIndex := 0; domItemsIndex < domItemsLen; domItemsIndex++ {
			respItem := new(Item)
			respItem.Mapping(dom[domItemsIndex])
			(*resp)[domItemsIndex] = respItem
		}
	}

	/**************** mapping end  ****************/
}
