package router

import (
	"github.com/gin-gonic/gin"
	"github.com/himanshuk42/stripe-charge/src/controller"
)

func ChargeRouter() *gin.Engine {
	r := gin.Default()
	g1 := r.Group("/stripe")
	{
		g1.POST("payment", controller.Payment)
	}
	return r
}
