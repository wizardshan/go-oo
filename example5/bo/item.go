package bo

const (
	ItemCategoryDiscount = iota + 1
	ItemCategoryTrial
	ItemCategoryRebate
)

type Items []*Item

type Item struct {
	ID          int
	Category    int
	Title       string
	Stock       int
	PriceMarket int
	Price       int
	Rebate      int
}

func (bo *Item) PriceCalculate() {
	if bo.Category == ItemCategoryDiscount { // 折扣商品价格
		bo.Price = bo.PriceMarket / 2
	} else if bo.Category == ItemCategoryTrial { // 试用商品价格
		bo.Price = 0
	} else if bo.Category == ItemCategoryRebate { // 返利商品价格
		bo.Price = bo.PriceMarket
	}
}

// 返利计算函数
func (bo *Item) RebateCalculate() {
	if bo.Category == ItemCategoryDiscount || bo.Category == ItemCategoryTrial { // 折扣商品和试用商品返利金额为零
		bo.Rebate = 0
	} else if bo.Category == ItemCategoryRebate { // 返利商品返利
		bo.Rebate = bo.PriceMarket * 5 / 100
	}
}
