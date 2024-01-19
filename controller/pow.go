package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kamuridesu/calculator-go/calculator"
)

func Pow(c *gin.Context) {
	var params Params
	if c.BindQuery(&params) != nil {
		c.JSON(400, gin.H{
			"error": "Reading parameters error",
		})
		return
	}
	c.JSON(200, calculator.Pow(params.First, params.Second))
}
