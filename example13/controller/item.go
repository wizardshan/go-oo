package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example13/repository"
	"go-oo/example13/response"
	"net/http"
)

type Item struct {
	repo     *repository.Item
	repoUser *repository.User
}

func NewItem() *Item {
	ctr := new(Item)
	ctr.repo = repository.NewItem()
	ctr.repoUser = repository.NewUser()

	return ctr
}

// 获取商品列表
func (ctr *Item) All(c *gin.Context) {
	user := ctr.repoUser.Get()
	items := ctr.repo.All()

	for _, item := range items {
		if user.IsLevelGold() {
			item.Price -= 1
		} else if user.IsLevelPlatinum() {
			item.Price -= 2
		} else if user.IsLevelDiamond() {
			item.Price -= 3
		}
	}

	resp := response.Items{}
	resp.Mapping(items)

	c.JSON(http.StatusOK, resp)
}

// 获取商品详情
func (ctr *Item) Get(c *gin.Context) {

	user := ctr.repoUser.Get()
	item := ctr.repo.Get()

	if user.IsLevelGold() {
		item.Price -= 1
	} else if user.IsLevelPlatinum() {
		item.Price -= 2
	} else if user.IsLevelDiamond() {
		item.Price -= 3
	}

	var resp response.Item
	resp.Mapping(item)

	c.JSON(http.StatusOK, resp)
}
