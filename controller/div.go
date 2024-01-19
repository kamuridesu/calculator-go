package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kamuridesu/calculator-go/calculator"
)

func Div(c *gin.Context) {
	var params Params
	if c.BindQuery(&params) != nil {
		c.JSON(400, gin.H{
			"error": "Reading parameters error",
		})
		return
	}
	c.JSON(200, calculator.Div(params.First, params.Second))
}
