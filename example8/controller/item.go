package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example8/domain"
	"go-oo/example8/repository"
	"go-oo/example8/response"
	"net/http"
)

type Item struct {
	repo *repository.Item
}

func NewItem() *Item {
	ctr := new(Item)
	ctr.repo = repository.NewItem()
	return ctr
}

func (ctr *Item) Get(c *gin.Context) {
	item := ctr.repo.Get()

	var resp response.Item
	resp.Mapping(item)

	c.JSON(http.StatusOK, resp)
}

func (ctr *Item) Order(c *gin.Context) {
	item := ctr.repo.Get()
	number := 2

	instance := item.OfInstance()
	// 断言是否为虚拟商品类型
	_, isItemVirtual := instance.(*domain.ItemVirtual)
	if !isItemVirtual && item.OutOfStock(number) {
		c.JSON(http.StatusOK, "库存不足")
	}
}
