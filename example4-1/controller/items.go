package controller

import (
	"github.com/gin-gonic/gin"
	"go-oo/example4/repository"
	"go-oo/example4/response"
	"net/http"
)

type Items struct {
	resp *repository.Item
}

func NewItems() *Items {
	ctr := new(Items)
	ctr.resp = repository.NewItem()
	return ctr
}

func (ctr *Items) Get(c *gin.Context) {
	items := ctr.resp.All()

	var resp response.Items
	resp.Mapping(items)

	c.JSON(http.StatusOK, resp)
}
