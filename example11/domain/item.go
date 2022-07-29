package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryRebate
	ItemCategoryTrial
)

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

type ItemPriceVIPCalculator interface {
	PriceVIP() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
}

// 市场价是否显示接口
type ItemPriceMarketHidden interface {
	PriceMarketHidden() bool
}

type Items []*Item

// 基类商品
type Item struct {
	ID       int
	Category int
	Title    string
	PriceMarket int
	Stock int

	Instance interface{}
}
