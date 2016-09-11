package routers

import (
	"github.com/NWPU-Paper/Paper/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	beego.Router("/teacher",&controllers.TeacherController{})
	beego.Router("/student",&controllers.StudentController{})
	beego.Router("/test",&controllers.TestController{})

	beego.Router("/subject/:id([0-9]+)", &controllers.SubjectController{}, "Get:Detail")

}
