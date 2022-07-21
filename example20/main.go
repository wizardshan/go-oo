package main

import (
	"github.com/gin-gonic/gin"
	"go-oo/example2/controller"
)

func main() {

	r := gin.Default()

	itemsCtr := controller.NewItems()
	r.GET("example2/items", itemsCtr.Get)

	r.Run()
}
