package router

import (
	"github.com/Mmx233/VpsBrokerC/middlewares"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	G = gin.New()
	G.Use(gin.Recovery(), middlewares.Auth())

	routerC2c(G.Group("/c"))
	routerC2p(G.Group("/p"))
}
