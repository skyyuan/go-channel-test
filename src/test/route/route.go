package route

import (
	"fmt"
	"sync"
	"test/utiles"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"test/models"
)

type UserConn struct {
	Conn chan models.Message
}

type Router struct {
	sync.RWMutex
	UserConns map[string]UserConn
}

func NewMessageRouter() *Router {
	r := &Router{UserConns: make(map[string]UserConn)}
	go routerSub(r)
	return r
}


func routerSub(r *Router) {
	conn := utiles.SubPool.Get()
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe("chat")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Println("Receive Message")
			r.dispatch(v.Data, v.Channel)
		case redis.Subscription:
			fmt.Println("订阅成功")
		case error:
			conn.Close()
			conn = utiles.SubPool.Get()
			psc = redis.PubSubConn{Conn: conn}
			psc.Subscribe("chat")

		}
	}
	conn.Close()
}


func (r *Router) dispatch(data []byte, channel string) {
	var m models.Message

	if err := json.Unmarshal(data, &m); err != nil {
		fmt.Println("model.Message Unmarshal error: ", err)
		return
	}
	if len(m.SignalID) > 0 {
		r.RLock()
		if c, ok := r.UserConns[m.SignalID]; ok {
			fmt.Println("push to ", m.SignalID)
			r.push(&c, &m)
		} else {
			fmt.Println("this user offline")
		}
		r.RUnlock()
	}
}

func (r *Router) push(c *UserConn, m *models.Message) {
	c.Conn <- *m
}

func (r *Router) ClearUserData(userID string) {
	r.RLock()
	_, ok := r.UserConns[userID]
	r.RUnlock()
	if ok {
		delete(r.UserConns, userID)
	}

}
