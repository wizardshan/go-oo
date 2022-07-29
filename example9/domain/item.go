package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
	ItemCategoryVirtual
)

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
}

// 市场价是否显示接口
type ItemPriceMarketHidden interface {
	PriceMarketHidden() bool
}

// 商品库存管理接口
type ItemStockHandler interface {
	OutOfStock(number int) bool
	Stock() int
}

type Items []*Item

// 基类商品
type Item struct {
	ID       int
	Category int
	Title    string
	PriceMarket int

	Instance interface{}
}

type ItemStock struct {
	Number       int
}

func (dom *ItemStock) OutOfStock(number int) bool {
	if dom.Number < number {
		return true
	}

	return false
}

func (dom *ItemStock) Stock() int {
	return dom.Number
}

