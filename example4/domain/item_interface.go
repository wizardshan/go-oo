package domain

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() *int
}

// 市场价是否显示接口
type ItemPriceMarketHidden interface {
	PriceMarketHidden() bool
}


