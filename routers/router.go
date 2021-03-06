package routers

import (
	"github.com/NWPU-Paper/Paper/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")

	teacherController := new(controllers.TeacherController)
	beego.Router("/teacher/index", teacherController)
	beego.Router("/teacher/profile", teacherController, "get:Profile")
	beego.Router("/teacher/subjects", teacherController, "get:Subject")
	beego.Router("/teacher/subjects/addsubject", teacherController, "get:AddSubject")

	beego.Router("/teacher/subjects/addSubject", teacherController, "post:ExecAddSubject")
	beego.Router("/teacher/paper", teacherController, "get:Paper")
	beego.Router("/teacher/score", teacherController, "get:Score")
	beego.Router("/teacher/assignments", teacherController, "get:Assignment")
	beego.Router("/teacher/students", teacherController, "get:Student")
	beego.Router("/teacher/message", teacherController, "get:Message")

	//studentController:=new(controllers.StudentController)
	//beego.Router("/student",studentController)
	//beego.Router("/student/profile",studentController,"get:Profile")

	beego.Router("/subject/:id([0-9]+)", &controllers.SubjectController{})
	beego.Router("/subject/:id([0-9]+)/lock/:user", &controllers.SubjectController{}, "get:Lock")
	beego.Router("/subject/:id([0-9]+)/select", &controllers.SubjectController{}, "get:Select")

	beego.Router("/list", &controllers.ListController{})
	beego.Router("/mylist", &controllers.ListController{}, "get:Select")

	beego.Router("/profile", &controllers.ProfileController{})

	beego.Router("/r", &controllers.BaseController{}, "post:R")

}
