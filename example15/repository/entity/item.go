package entity

import (
	"go-oo/example15/bo"
)

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
}

func (ent *Item) Mapping() *bo.Item {
	boItem := new(bo.Item)
	boItem.ID = ent.ID
	boItem.Category = ent.Category
	boItem.Title = ent.Title
	boItem.Stock = ent.Stock
	boItem.PriceMarket = ent.PriceMarket

	boItem.OfInstance()
	// 断言计算价格
	if priceCalculator, ok := boItem.Instance.(bo.ItemPriceCalculator); ok {
		boItem.Price = priceCalculator.Price()
	}
	// 断言计算VIP价格
	if priceVIPCalculator, ok := boItem.Instance.(bo.ItemPriceVIPCalculator); ok {
		boItem.PriceVIP = priceVIPCalculator.PriceVIP()
	}

	// 断言计算返利
	if rebateCalculator, ok := boItem.Instance.(bo.ItemRebateCalculator); ok {
		boItem.Rebate = rebateCalculator.Rebate()
	}

	return boItem
}

func (ents Items) Mapping() bo.Items {
	entsLen := len(ents)
	bos := make(bo.Items, entsLen)
	if entsLen > 0 {
		for entIndex := 0; entIndex < entsLen; entIndex++ {
			bos[entIndex] = ents[entIndex].Mapping()
		}
		return bos
	}

	return nil
}
