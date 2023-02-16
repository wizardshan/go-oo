package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example4/repository"
	"go-oo/example4/response"
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

	resp := response.Items{}
	resp.Mapping(items)

	c.JSON(http.StatusOK, resp)
}

// 获取商品详情
func (ctr *Item) Get(c *gin.Context) {
	item := ctr.repo.Get()

	var resp response.Item
	resp.Mapping(item)

	c.JSON(http.StatusOK, resp)
}
