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
	c.Data["pageKey"] = "index";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Profile() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/profile.tpl"
	c.LayoutSections = make(map[string]string)

	c.Data["pageKey"] = "profile";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Subject() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/subject.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "subject";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Paper() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/paper.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "paper";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Score() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/score.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "score";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Assignment() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/assignment.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "assignment";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Student() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/student.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "student";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}

func (c *TeacherController) Message() {
	c.Layout = "base/base.tpl"
	c.TplName = "teacher/message.tpl"

	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "message";
	c.LayoutSections["HtmlHead"] = "teacher/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "teacher/script.tpl"
}