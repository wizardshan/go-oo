package bo

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

// VIP价格计算器接口
type ItemPriceVIPCalculator interface {
	PriceVIP() int
}

// 返利计算器接口
type ItemRebateCalculator interface {
	Rebate() int
}

// 市场价是否隐藏接口
type ItemPriceMarketHidden interface {
	PriceMarketHidden() bool
}

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
	PriceVIP    int
	Rebate      *int

	Instance interface{}

	Shop *Shop
}

func (bo *Item) Of() {
	bo.instance()

	bo.priceCalculate()
	bo.priceVipCalculate()
	bo.rebateCalculate()
}

func (bo *Item) priceCalculate() {
	// 断言计算价格
	if priceCalculator, ok := bo.Instance.(ItemPriceCalculator); ok {
		bo.Price = priceCalculator.Price()
	}

	// 根据店铺等级计算价格
	if bo.Shop.IsLevelBronzeMedal() {
		bo.Price -= 1
	} else if bo.Shop.IsLevelSilverMedal() {
		bo.Price -= 2
	} else if bo.Shop.IsLevelGoldMedal() {
		bo.Price -= 3
	}
}

func (bo *Item) priceVipCalculate() {
	// 断言计算VIP价格
	if priceVIPCalculator, ok := bo.Instance.(ItemPriceVIPCalculator); ok {
		bo.PriceVIP = priceVIPCalculator.PriceVIP()
	}
}

func (bo *Item) rebateCalculate() {
	// 断言计算返利
	if rebateCalculator, ok := bo.Instance.(ItemRebateCalculator); ok {
		rebate := rebateCalculator.Rebate()
		bo.Rebate = &rebate
	}
}

func (bo *Item) instance() {

	var instance interface{}

	switch bo.Category {
	case ItemCategoryDiscount:
		instance = &ItemDiscount{
			Item: bo,
		}
	case ItemCategoryRebate:
		instance = &ItemRebate{
			Item: bo,
		}
	}

	bo.Instance = instance
}
