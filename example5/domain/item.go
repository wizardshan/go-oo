package domain

// 价格计算器接口
type ItemPriceCalculator interface {
	Price() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
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
	case 1:
		return &ItemDiscount{
			Item: dom,
		}
	case 2:
		return &ItemTrial{
			Item: dom,
		}
	case 3:
		return &ItemRebate{
			Item: dom,
		}
	}

	return nil
}