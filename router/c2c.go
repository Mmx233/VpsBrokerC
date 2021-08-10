package router

import (
	controllers "github.com/Mmx233/VpsBrokerC/controllers/c2c"
	"github.com/gin-gonic/gin"
)

func routerC2c(G *gin.RouterGroup) {
	G.GET("/", controllers.ConnectWs)
}
