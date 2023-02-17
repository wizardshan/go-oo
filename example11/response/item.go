package response

import (
	"go-oo/example11/bo"
)

type Items []*Item

type Item struct {
	ID                int    `json:"id"`
	Category          int    `json:"category"`
	Title             string `json:"title"`
	Stock             int    `json:"stock"`
	PriceMarket       int    `json:"priceMarket"`
	Price             int    `json:"price"`
	PriceVIP          *int   `json:"priceVIP,omitempty"`
	Rebate            int    `json:"rebate,omitempty"`
	PriceMarketHidden bool   `json:"priceMarketHidden"`
}

func (resp *Item) Mapping(boItem *bo.Item) {
	if boItem == nil {
		return
	}

	resp.ID = boItem.ID
	resp.Category = boItem.Category
	resp.Title = boItem.Title
	resp.Stock = boItem.Stock
	resp.PriceMarket = boItem.PriceMarket
	resp.Price = boItem.Price
	resp.Rebate = boItem.Rebate

	// 断言市场价是否显示
	if rebateCalculator, ok := boItem.Instance.(bo.ItemPriceMarketHidden); ok {
		resp.PriceMarketHidden = rebateCalculator.PriceMarketHidden()
	}

	if priceCalculator, ok := boItem.Instance.(bo.ItemPriceCalculator); ok {
		resp.Price = priceCalculator.Price()

		// 出现面向实现编程的问题
		if _, ok := boItem.Instance.(bo.ItemTrial); !ok {
			priceVIP := priceCalculator.PriceVIP()
			resp.PriceVIP = &priceVIP
		}
	}
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
