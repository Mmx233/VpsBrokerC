package controllers

import (
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upper = websocket.Upgrader{
	HandshakeTimeout: time.Minute,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

	conn, err := upper.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// bad request
		return
	}

	go Conn.MakeConnChan(ip, n.Port, conn)
}
