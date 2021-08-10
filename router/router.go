package router

import (
	"github.com/Mmx233/VpsBrokerC/middlewares"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	G = gin.Default()
	G.Use(middlewares.Auth())

	routerC2c(G.Group("/c"))
	routerC2p(G.Group("/p"))
}
