package controllers

import (
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

}
