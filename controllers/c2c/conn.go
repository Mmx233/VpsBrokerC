package controllers

import (
	"fmt"
	"github.com/Mmx233/VpsBrokerC/global"
	"github.com/Mmx233/VpsBrokerC/models"
	"github.com/Mmx233/VpsBrokerC/service"
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
func (a *conn) Connect(ip string, port uint) (*websocket.Conn,error) {
	a.lock.RLock()
	conn, ok := a.Pool[ip]
	a.lock.RUnlock()
	if ok {
		_ = conn.Close()
	}

	var e error
	a.lock.Lock()
	defer a.lock.Unlock()
	a.Pool[ip], _, e = websocket.DefaultDialer.Dial(
		"ws://"+ip+":"+fmt.Sprint(port), nil)

	return a.Pool[ip],e
}

// Renew 与neighbor列表同步
func (a *conn) Renew() {
	global.Neighbors.Lock.RLock()
	defer global.Neighbors.Lock.RUnlock()
	a.lock.Lock()
	defer a.lock.Unlock()

	//断开被删除客户端
	var del bool
	for k:=range a.Pool {
		if _ ,ok:=global.Neighbors.Data[k];!ok {
			_ = a.Pool[k].Close()
			delete(a.Pool,k)
			del=true
		}
	}

	//连接新客户端
	for k,v:=range global.Neighbors.Data {
		if _,ok:=a.Pool[k];!ok {
			go a.Connection(k,v.Port)
		}
	}

	//回收内存
	if del {
		var t =make(map[string]*websocket.Conn,len(a.Pool))
		for k,v:=range a.Pool {
			t[k]=v
		}
		a.Pool=t
	}
}

// Connection 客户端连接协程
func (a *conn)Connection(ip string,port uint){
	start:
		//建立连接
	conn,e:=a.Connect(ip,port)
	if e!=nil {
		for e!=nil {
			log.Println("连接client失败：/n",e)
			time.Sleep(time.Second)
			conn,e=a.Connect(ip,port)
		}
	}

	//心跳
	var hb models.HeartBeat
	var c = make(chan error,1)
	go func() {//接收心跳包
		for e!=nil {
			e=conn.ReadJSON(&hb)
			c<-e
		}
	}()
	go func() {//发送心跳包
		for conn.WriteJSON(&models.HeartBeat{
			Type: "heartbeat",
			Time: time.Now().UnixNano(),
		})!=nil {
			time.Sleep(time.Second)
		}
	}()
	for  {//处理心跳超时
		select {
		case <-c:
			if e==nil {
				//todo up
			}
		case <-time.After(time.Second*5):
			service.Msg.Channel<-&models.HeartBeat{
			Type: "down",
			TargetIp: ip,
			Time: hb.Time,
			}
		}
		if e!=nil {
			break
		}
	}

	//连接断开
	a.lock.RLock()
	if _,ok:=a.Pool[ip];!ok {
		conn=nil
		a.lock.RUnlock()
		return
	}else {
		a.lock.RUnlock()
		goto start
	}
}
