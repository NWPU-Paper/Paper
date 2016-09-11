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

	if app, ok := c.AppController.(LoginPreparer); ok {
		app.LoginPrepare()
	}

	c.Data["CurrentUser"] = c.user

	c.Layout = "base/base.tpl"

	c.LayoutSections = make(map[string]string)

	c.LayoutSections["HtmlHead"] = "base/htmlhead.tpl"

	switch c.user.Type {
	case models.TYPE_USER_STUDENT:
		break
	case models.TYPE_USER_TEACHER:
		c.Data["Module"] = "老师"
		c.LayoutSections["Nav"] = "teacher/header.tpl"
		break
	case models.TYPE_USER_LEADER:
		c.Data["Module"] = "专业负责人"

		break
	case models.TYPE_USER_SECRETARY:
		c.Data["Module"] = "答辩秘书"
		c.LayoutSections["Nav"] = "secretar/header.tpl"
		break
	case models.TYPE_USER_ADMIN:
		c.Data["Module"] = "管理员"

		break
	default:
		c.Abort("403")
	}

	c.Data["Title"] = "控制面板"
	
}

// 管理的首页
func (c *AdminController) Get() {
	c.TplName = "index.tpl"

}

