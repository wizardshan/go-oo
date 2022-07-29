package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example9/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example9/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example9/item", itemCtr.Get)
	r.GET("example9/item/order", itemCtr.Order)

	r.Run()
}
