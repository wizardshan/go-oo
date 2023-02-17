package entity

import (
	"go-oo/example5/bo"
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
	bo := new(bo.Item)
	bo.ID = ent.ID
	bo.Category = ent.Category
	bo.Title = ent.Title
	bo.Stock = ent.Stock
	bo.PriceMarket = ent.PriceMarket
	bo.PriceCalculate()  // 调用业务模型函数Price，计算商品价格
	bo.RebateCalculate() // 调用业务模型函数Rebate，计算返利金额
	return bo
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
