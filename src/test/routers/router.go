package routers

import (
	"test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("test/",
			beego.NSInclude(
				&controllers.MainController{},
			),
	)

	beego.AddNamespace(ns)

}
