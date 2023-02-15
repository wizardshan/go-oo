package response

import (
	"go-oo/example1/bo"
)

type Items []*Item

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Stock       int    `json:"stock"`
	PriceMarket int    `json:"priceMarket"`
	Price       int    `json:"price"`
}

func (resp *Item) Mapping(bo *bo.Item) {
	if bo == nil {
		return
	}

	resp.ID = bo.ID
	resp.Title = bo.Title
	resp.Stock = bo.Stock
	resp.PriceMarket = bo.PriceMarket
	resp.Price = bo.Price
}

func (resp *Items) Mapping(bos bo.Items) {
	bosLen := len(bos)
	*resp = make(Items, bosLen)
	if bosLen > 0 {
		for boIndex := 0; boIndex < bosLen; boIndex++ {
			respItem := new(Item)
			respItem.Mapping(bos[boIndex])
			(*resp)[boIndex] = respItem
		}
	}
}
