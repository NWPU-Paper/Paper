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
	c.Data["Title"] = "老师面板"
}

func (c *TeacherController) Get() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/index.tpl"

	c.LayoutSections = make(map[string]string)

	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/header.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}