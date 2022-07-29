package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example11/repository"
	"go-oo/example11/response"
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