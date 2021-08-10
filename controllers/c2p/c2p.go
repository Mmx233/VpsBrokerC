package controllers

import (
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gin-gonic/gin"
	"time"
)

func PanelConnection(c *gin.Context) {
	conn, e := util.Upper.Upgrade(c.Writer, c.Request, nil)
	if e != nil {
		return
	}

	for conn.WriteJSON(models.Stat{}.Gen()) == nil {
		time.Sleep(time.Second)
	}
}
