package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type conn struct {
	Pool map[string]*websocket.Conn
	lock *sync.RWMutex
}

var Conn = conn{
	Pool: make(map[string]*websocket.Conn),
	lock: &sync.RWMutex{},
}

func (a *conn) Connect(ip string, port uint) error {
	a.lock.RLock()
	conn, ok := a.Pool[ip]
	a.lock.RUnlock()
	if ok {
		_ = conn.Close()
	}

	var e error
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.Pool[ip], _, e = websocket.DefaultDialer.Dial(
		"ws://"+ip+":"+fmt.Sprint(port), nil); e != nil {
		return e
	}

	return nil
}

func (*conn) Renew() {

}
