package bo

// 价格计算器接口
type ItemPriceCalculator interface {
	PriceCalculate()
}

// 返利计算器接口
type ItemRebateCalculator interface {
	RebateCalculate() int
}

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int

	Instance interface{}
}

func (bo *Item) Of() {
	bo.instance()
	// 断言计算价格
	if priceCalculator, ok := bo.Instance.(ItemPriceCalculator); ok {
		priceCalculator.PriceCalculate()
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
