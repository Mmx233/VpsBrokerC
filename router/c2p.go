package router

import (
	controllers "github.com/Mmx233/VpsBrokerC/controllers/c2p"
	"github.com/gin-gonic/gin"
)

func routerC2p(G *gin.RouterGroup) {
	G.GET("/", controllers.PanelConnection)
}
