package response

import "go-oo/example9/domain"

type Items []*Item

type Item struct {
	ID          int    `json:"id"`
	Category    int    `json:"category"`
	Title       string `json:"title"`
	Stock       *int    `json:"stock,omitempty"`
	PriceMarket int    `json:"priceMarket"`
	PriceMarketHidden bool   `json:"priceMarketHidden"`
	Price       int    `json:"price"`
	Rebate      *int   `json:"rebate,omitempty"`
}

func (resp *Item) Mapping(dom *domain.Item) {
	resp.ID = dom.ID
	resp.Category = dom.Category
	resp.Title = dom.Title
	resp.PriceMarket = dom.PriceMarket

	// 断言计算价格
	if priceCalculator, ok := dom.Instance.(domain.ItemPriceCalculator); ok {
		resp.Price = priceCalculator.Price()
	}

	// 断言计算返利
	if rebateCalculator, ok := dom.Instance.(domain.ItemRebateCalculator); ok {
		rebate := rebateCalculator.Rebate()
		resp.Rebate = &rebate
	}

	// 断言市场价是否显示
	if rebateCalculator, ok := dom.Instance.(domain.ItemPriceMarketHidden); ok {
		resp.PriceMarketHidden = rebateCalculator.PriceMarketHidden()
	}

	// 断言获取库存
	if stockHandler, ok := dom.Instance.(domain.ItemStockHandler); ok {
		stock := stockHandler.Stock()
		resp.Stock = &stock
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
