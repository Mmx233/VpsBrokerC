package controllers

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gin-gonic/gin"
)

func ConnectWs(c *gin.Context) {
	ip := c.ClientIP()

	global.Neighbors.Lock.RLock()
	defer global.Neighbors.Lock.RUnlock()
	n, ok := global.Neighbors.Data[ip]
	if !ok {
		//不在列表中
		return
	}

	if Conn.Connected(ip) {
		//已经连上了
		return
	}

	conn, err := util.Upper.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// bad request
		return
	}

	go Conn.MakeConnChan(ip, n.Port, conn)
}
