package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
)

type TeacherController struct {
	AdminController
}

func (c *TeacherController) LoginPrepare() {

	if c.user.Type != models.TYPE_USER_TEACHER {
		c.Abort("403")
	}
}

func (c *TeacherController) Get() {
	//c.Layout = "base/base.tpl"
	//c.TplName = "teacher/index.tpl"
	//c.LayoutSections["HtmlHead"] = "footer.tpl"
	//c.LayoutSections["SideBar"] = "footer.tpl"
	//c.LayoutSections["Nav"] = "footer.tpl"
	c.Ctx.WriteString("这是老师界面")
}