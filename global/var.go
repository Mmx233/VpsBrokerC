package global

import "github.com/gorilla/websocket"

var Conn *websocket.Conn

var Neighbors map[string]uint
