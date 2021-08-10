package router

import (
	controllers "github.com/Mmx233/VpsBrokerC/controllers/c2c"
	"github.com/Mmx233/VpsBrokerC/middlewares"
	"github.com/gin-gonic/gin"
)

var G *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	G = gin.Default()
	G.Use(middlewares.Auth())

	G.GET("/", controllers.ConnectWs)
}
