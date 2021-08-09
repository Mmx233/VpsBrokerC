package controllers

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/global"
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

func (a *conn) Renew() {
	global.Neighbors.Lock.RLock()
	defer global.Neighbors.Lock.RUnlock()
	a.lock.Lock()
	defer a.lock.Unlock()

	var del bool
	for k,_:=range a.Pool {
		if _ ,ok:=global.Neighbors.Data[k];!ok {
			_ = a.Pool[k].Close()
			delete(a.Pool,k)
			del=true
		}
	}

	for k,v:=range global.Neighbors.Data {
		if _,ok:=a.Pool[k];!ok {
			go a.Connection(k,v)
		}
	}

	if del {
		var t =make(map[string]*websocket.Conn,len(a.Pool))
		for k,v:=range a.Pool {
			t[k]=v
		}
		a.Pool=t
	}
}

func (*conn)Connection(ip string,port uint){

	//todo 断连时检查是否是主动删除
}
