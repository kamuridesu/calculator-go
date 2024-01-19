package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kamuridesu/calculator-go/controller"
)

func main() {
	r := gin.Default()
	r.GET("/add", controller.Add)
	r.GET("/sub", controller.Sub)
	r.GET("/mul", controller.Mul)
	r.GET("/div", controller.Div)
	r.GET("/pow", controller.Pow)
	r.Run()
}
