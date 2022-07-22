package response

import "go-oo/example3/domain"

type Items []*Item

type Item struct {
	ID                int    `json:"id"`
	Category          int    `json:"category"`
	Title             string `json:"title"`
	Stock             int    `json:"stock"`
	PriceMarket       int    `json:"priceMarket"`
	PriceMarketHidden bool   `json:"priceMarketHidden"`
	Price             int    `json:"price"`
	Rebate            *int   `json:"rebate,omitempty"`
}

func (resp *Item) Mapping(dom *domain.Item) {
	resp.ID = dom.ID
	resp.Category = dom.Category
	resp.Title = dom.Title
	resp.Stock = dom.Stock
	resp.PriceMarket = dom.PriceMarket
	resp.Price = dom.Price()
	resp.Rebate = dom.Rebate()
	resp.PriceMarketHidden = dom.PriceMarketHidden()
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
