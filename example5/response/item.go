package response

import "go-oo/example5/domain"

type Items []*Item

type Item struct {
	ID          int    `json:"id"`
	Category    int    `json:"category"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	PriceMarket int    `json:"priceMarket"`
	Price       int    `json:"price"`
	Rebate      int    `json:"rebate"`
}

func (resp *Item) Mapping(dom *domain.Item) {
	resp.ID = dom.ID
	resp.Category = dom.Category
	resp.Title = dom.Title
	resp.Stock = dom.Stock
	resp.PriceMarket = dom.PriceMarket

	instance := dom.OfInstance()
	// 断言计算价格
	if priceCalculator, ok := instance.(domain.ItemPriceCalculator); ok {
		resp.Price = priceCalculator.Price()
	}

	// 断言计算返利
	if rebateCalculator, ok := instance.(domain.ItemRebateCalculator); ok {
		resp.Rebate = rebateCalculator.Rebate()
	}
}

func (resp *Items) Mapping(dom domain.Items) {
	domItemsLen := len(dom)
	*resp = make(Items, domItemsLen)
	if domItemsLen > 0 {
		for domItemsIndex := 0; domItemsIndex < domItemsLen; domItemsIndex++ {
			respItem := new(Item)
			respItem.Mapping(dom[domItemsIndex])
			(*resp)[domItemsIndex] = respItem
		}
	}
}
