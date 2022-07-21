package response

import "go-oo/example1/domain"

type Items []*Item

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	PriceMarket int    `json:"priceMarket"`
	Price       int    `json:"price"`
}

func (resp *Item) Mapping(dom *domain.Item) {
	resp.ID = dom.ID
	resp.Title = dom.Title
	resp.Stock = dom.Stock
	resp.PriceMarket = dom.PriceMarket
	resp.Price = dom.Price()       // 调用业务模型函数Price，计算商品价格
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
