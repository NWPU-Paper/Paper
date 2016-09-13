package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
)

type LoginPreparer interface{
	LoginPrepare()
}


//所有需要登陆的都继承这个 未登陆则自动返回
type AdminController struct {
	BaseController
}

func (c *AdminController) NestPrepare() {
	if !c.isLogin {
		//未登录则返回401
		c.RedirectTo("LoginController.Get")
	}



	c.Data["CurrentUser"] = c.user

	c.Layout = "base/base2.tpl"

	c.LayoutSections = make(map[string]string)

	c.LayoutSections["HtmlHead"] = "base/htmlhead.tpl"

	c.Data["pageKey"] = ""

	switch c.user.Type {
	case models.TYPE_USER_STUDENT:
		c.Data["Module"] = "Student"
		c.LayoutSections["Nav"] = "student/nav.tpl"
		break
	case models.TYPE_USER_TEACHER:
		c.Data["Module"] = "老师"
		c.LayoutSections["Nav"] = "teacher/nav.tpl"
		break
	case models.TYPE_USER_LEADER:
		c.Data["Module"] = "专业负责人"
		c.LayoutSections["Nav"] = "leader/nav.tpl"
		break
	case models.TYPE_USER_SECRETARY:
		c.Data["Module"] = "答辩秘书"
		c.LayoutSections["Nav"] = "secretary/nav.tpl"
		break
	case models.TYPE_USER_ADMIN:
		c.Data["Module"] = "管理员"

		break
	default:
		c.Abort("403")
	}

	c.Data["Title"] = "控制面板"

	if app, ok := c.AppController.(LoginPreparer); ok {
		app.LoginPrepare()
	}
	
}

// 管理的首页
func (c *AdminController) Get() {
	c.TplName = "index.tpl"

}

