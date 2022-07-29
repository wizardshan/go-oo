package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example8/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example8/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example8/item", itemCtr.Get)
	r.GET("example8/item/order", itemCtr.Order)

	r.Run()
}
