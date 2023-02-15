package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example1/repository"
	"go-oo/example1/response"
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

// 获取商品列表
func (ctr *Item) All(c *gin.Context) {
	items := ctr.repo.All()

	for _, item := range items {
		item.Price = item.PriceMarket / 2 // 折扣计算逻辑
	}

	resp := response.Items{}
	resp.Mapping(items)

	c.JSON(http.StatusOK, resp)
}

// 获取商品详情
func (ctr *Item) Get(c *gin.Context) {
	item := ctr.repo.Get()
	item.Price = item.PriceMarket / 2 // 折扣计算逻辑

	var resp response.Item
	resp.Mapping(item)

	c.JSON(http.StatusOK, resp)
}
