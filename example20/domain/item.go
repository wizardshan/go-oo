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

const (
	ItemCategoryRebate = iota + 1
	ItemCategoryDiscount
)

type Items []*Item

// 基类商品
type Item struct {
	ID       int
	Category int
	Title    string
	Stock    int

	PriceMarket int
}

func (dom *Item) OfInstance() interface{} {
	switch dom.Category {
	case ItemCategoryRebate:
		return dom.OfInstanceRebate()
	case ItemCategoryDiscount:
		return dom.OfInstanceDiscount()
	}

	return nil
}

func (dom *Item) OfInstanceRebate() *ItemRebate {
	return &ItemRebate{
		Item: dom,
	}
}

func (dom *Item) OfInstanceDiscount() *ItemDiscount {
	return &ItemDiscount{
		Item: dom,
	}
}

func (dom *Item) StockEnough(number int) bool {
	if dom.Stock >= number {
		return true
	}

	return false
}
