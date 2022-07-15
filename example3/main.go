package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example3/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example3/items", itemsCtr.Get)

	itemCtr := controller.NewItem()
	r.GET("example3/item", itemCtr.Get)
	r.Run()
}
