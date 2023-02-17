package domain

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
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
	case ItemCategoryDiscount:
		return &ItemDiscount{
			Item: dom,
		}
	case ItemCategoryTrial:
		return &ItemTrial{
			Item: dom,
		}
	case ItemCategoryRebate:
		return &ItemRebate{
			Item: dom,
		}
	}

	return nil
}

func (dom *Item) OutOfStock(number int) bool {
	if dom.Stock < number {
		return true
	}

	return false
}