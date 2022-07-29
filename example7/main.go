package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example7/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example7/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example7/item", itemCtr.Get)
	r.GET("example7/item/order", itemCtr.Order)

	r.Run()
}
