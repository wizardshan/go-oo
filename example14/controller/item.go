package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example14/repository"
	"go-oo/example14/response"
	"go-oo/example14/service"
	"net/http"
)

type Item struct {
	repo     *repository.Item
	repoUser *repository.User

	prizeCalculator *service.ItemPrizeCalculator
}

func NewItem() *Item {
	ctr := new(Item)
	ctr.repo = repository.NewItem()
	ctr.repoUser = repository.NewUser()
	ctr.prizeCalculator = service.NewItemPrizeCalculator()

	return ctr
}

// 获取商品列表
func (ctr *Item) All(c *gin.Context) {
	user := ctr.repoUser.Get()
	items := ctr.repo.All()

	for _, item := range items {
		ctr.prizeCalculator.Handle(item, user)
	}

	resp := response.Items{}
	resp.Mapping(items)

	c.JSON(http.StatusOK, resp)
}

// 获取商品详情
func (ctr *Item) Get(c *gin.Context) {

	user := ctr.repoUser.Get()
	item := ctr.repo.Get()
	ctr.prizeCalculator.Handle(item, user)

	var resp response.Item
	resp.Mapping(item)

	c.JSON(http.StatusOK, resp)
}
