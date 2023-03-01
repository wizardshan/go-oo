package bo

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
)

// 价格计算器接口
type ItemPriceCalculator interface {
	PriceCalculate()
}

// 返利计算器接口
type ItemRebateCalculator interface {
	RebateCalculate()
}

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
	Rebate      int

	Instance interface{}
}

func (bo *Item) Of() {
	bo.instance()
	// 断言计算价格
	if priceCalculator, ok := bo.Instance.(ItemPriceCalculator); ok {
		priceCalculator.PriceCalculate()
	}

	// 断言计算返利
	if rebateCalculator, ok := bo.Instance.(ItemRebateCalculator); ok {
		rebateCalculator.RebateCalculate()
	}
}

func (bo *Item) instance() {

	var instance interface{}
	switch bo.Category {
	case 1:
		instance = &ItemDiscount{
			Item: bo,
		}
	case 2:
		instance = &ItemTrial{
			Item: bo,
		}
	case 3:
		instance = &ItemRebate{
			Item: bo,
		}
	}

	bo.Instance = instance
}
