package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"test/route"
	"test/models"
	"test/utiles"
)

type MainController struct {
	beego.Controller
}



// @router / [get]
func (c *MainController) Get() {
	messageRouter := route.NewMessageRouter()
	signalId := c.GetString("a")
	ch := make(chan models.Message)

	defer messageRouter.ClearUserData(signalId)
	messageRouter.Lock()
	messageRouter.UserConns[signalId] = route.UserConn{Conn: ch}
	a, ok := messageRouter.UserConns[signalId];
	fmt.Println(a)
	fmt.Println(ok)
	messageRouter.Unlock()
	timeout := time.After(300 * time.Second)
	var m models.Message
	select {
	case fk := <- ch:
	// get data from ch
		fmt.Println("1")
		fmt.Println(fk)
		m = fk

	case <- timeout:
	// read data from ch timeout
		fmt.Println("超时")
	}

	c.Data["json"] = m
	c.ServeJSON()
}
// @router /list [get]
func (c *MainController) List() {
	m := c.GetString("m")
	conn := utiles.SubPool.Get()
	conn.Do("PUBLISH", "chat", m)

	c.Data["json"] = map[string]bool{"success": true}
	c.ServeJSON()
}

