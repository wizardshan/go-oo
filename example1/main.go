package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example1/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example1/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example1/item", itemCtr.Get)
	r.Run()
}
