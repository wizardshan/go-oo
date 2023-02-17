package bo

// 商品库存管理接口
type ItemStockHandler interface {
	OutOfStock(number int) bool
}

type ItemStock struct {
	Number int
}

func (bo *ItemStock) OutOfStock(number int) bool {
	if bo.Number < number {
		return true
	}

	return false
}
