package routers

import (
	"github.com/banixc/paper/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{})
	beego.Router("/logout",&controllers.AuthController{},"*:Logout")
	beego.Router("/list",&controllers.ListController{})

}
