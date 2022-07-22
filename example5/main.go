package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example5/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example5/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example5/item", itemCtr.Get)
	r.Run()
}
