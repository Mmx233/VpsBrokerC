package controllers

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/service"
	"github.com/Mmx233/VpsBrokerC/util"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

type conn struct {
	Pool map[string]*websocket.Conn
	lock *sync.RWMutex
}

// Conn c2c连接池
var Conn = conn{
	Pool: make(map[string]*websocket.Conn),
	lock: &sync.RWMutex{},
}

// Connect ws连接
func (a *conn) Connect(ip string, port uint) (*websocket.Conn, error) {
	a.lock.Lock()
	defer a.lock.Unlock()
	conn, ok := a.Pool[ip]
	if ok {
		_ = conn.Close()
	}

	conn2, _, e := websocket.DefaultDialer.Dial(
		"ws://"+ip+":"+fmt.Sprint(port)+"/c/", global.AuthHeader)

	return conn2, e
}

func (a *conn) Connected(ip string) bool {
	a.lock.RLock()
	v, ok := a.Pool[ip]
	a.lock.RUnlock()
	return ok && v != nil
}

// Renew 与neighbor列表同步
func (a *conn) Renew() {
	global.Neighbors.Lock.RLock()
	defer global.Neighbors.Lock.RUnlock()
	a.lock.Lock()
	defer a.lock.Unlock()

	//断开被删除客户端
	var del bool
	for k := range a.Pool {
		if _, ok := global.Neighbors.Data[k]; !ok {
			_ = a.Pool[k].Close()
			delete(a.Pool, k)
			del = true
		}
	}

	//连接新客户端
	for k, v := range global.Neighbors.Data {
		if _, ok := a.Pool[k]; !ok {
			go a.ForceConnect(k, v.Port)
		}
	}

	//回收内存
	if del {
		var t = make(map[string]*websocket.Conn, len(a.Pool))
		for k, v := range a.Pool {
			t[k] = v
		}
		a.Pool = t
	}
}

// MakeConnChan 处理客户端连接协程
func (a *conn) MakeConnChan(ip string, port uint, conn *websocket.Conn) {
	a.lock.Lock()
	a.Pool[ip] = conn
	a.lock.Unlock()
	util.Event.NewPeer(ip)

	//心跳
	var e error
	var hb models.HeartBeat
	var c = make(chan error, 1)
	go func() { //接收心跳包
		for e == nil {
			e = conn.ReadJSON(&hb)
			c <- e
		}
	}()
	go func() { //发送心跳包
		for conn.WriteJSON(&models.HeartBeat{
			Type: "heartbeat",
			Time: time.Now().UnixNano(),
		}) == nil {
			time.Sleep(time.Second)
		}
	}()
	for e == nil { //处理心跳超时
		select {
		case <-c:
			if e == nil {
				service.Stat.Up(ip, hb.Time)
			}
		case <-time.After(time.Second * 5):
			service.Stat.Down(ip)
		}
	}

	//连接断开
	a.lock.RLock()
	defer a.lock.RUnlock()
	if conn == nil { //预防冲突
		return
	}
	_ = conn.Close()
	conn = nil
	util.Event.LostPeer(ip)
	if _, ok := a.Pool[ip]; !ok {
		return
	}
	go a.ForceConnect(ip, port)
}

// ForceConnect 主动连接客户端
func (a *conn) ForceConnect(ip string, port uint) {
	conn, e := a.Connect(ip, port)
	if e != nil {
		for e != nil {
			log.Println("连接client "+ip+" 失败：\n", e)
			time.Sleep(time.Second)
			conn, e = a.Connect(ip, port)
		}
	}

	a.MakeConnChan(ip, port, conn)
}
