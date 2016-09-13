package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
	"time"
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
	c.LayoutSections["Nav"] = "nav.tpl"

	c.TplName = "index.tpl"
	c.Data["pageKey"] = "index"
	Start := models.GetDate(0)
	c.Data["startDate"] = Start
	End := models.GetDate(1)
	c.Data["endDate"] = End
	c.Data["nowDate"] = time.Now()
	s := End.Date.Day() - time.Now().Day()
	if s < 0 {
		s = 0
	}
	c.Data["endDay"] = s
	switch c.user.Type {
	case models.TYPE_USER_STUDENT:
		c.Data["Module"] = "学生"
		break
	case models.TYPE_USER_TEACHER:
		c.Data["Module"] = "老师"
		break
	case models.TYPE_USER_LEADER:
		c.Data["Module"] = "专业负责人"
		break
	case models.TYPE_USER_SECRETARY:
		c.Data["Module"] = "答辩秘书"
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

}

// 管理编辑
func (c *AdminController) Post() {
	EndDate := c.GetString("end_date")
	p, err := time.Parse("2006-01-02",EndDate)
	if err == nil {
		models.SetDate(1, p)
	}
	c.RedirectTo("AdminController.Get")
}

