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
	c.Data["Title"] = "控制面板"

}

func (c *TeacherController) Get() {
	c.Data["pageKey"] = "index";

	c.TplName = "teacher/index.tpl"

}

func (c *TeacherController) Profile() {
	c.TplName = "teacher/profile.tpl"

	c.Data["pageKey"] = "profile";

}

func (c *TeacherController) Subject() {
	c.TplName = "teacher/subject.tpl"
	c.Data["pageKey"] = "subject";

}

func (c *TeacherController) Paper() {
	c.TplName = "teacher/paper.tpl"

	c.Data["pageKey"] = "paper";

}

func (c *TeacherController) Score() {
	c.TplName = "teacher/score.tpl"

	c.Data["pageKey"] = "score";
}

func (c *TeacherController) Assignment() {
	c.TplName = "teacher/assignment.tpl"
	c.Data["pageKey"] = "assignment";

}

func (c *TeacherController) Student() {
	c.TplName = "teacher/student.tpl"
	c.Data["pageKey"] = "student";

}

func (c *TeacherController) Message() {
	c.TplName = "teacher/message.tpl"
	c.Data["pageKey"] = "message";

}

// 发布新题目
func (c *TeacherController) AddSubject() {
	c.TplName = "teacher/addNewSubject.tpl"
	c.LayoutSections = make(map[string]string)
	c.Data["pageKey"] = "subject";
	c.LayoutSections["HtmlHead"] = "base/htmlhead.tpl"
	c.LayoutSections["Nav"] = "teacher/nav.tpl"
	c.LayoutSections["Script"] = "base/script.tpl"
}

// 处理发布新题目请求
func (c *TeacherController) ExecAddSubject() {
	title:=c.GetString("title")
	introduce:=c.GetString("introduce")
	status:=models.Status{Id:10}
	newSubject:=models.Subject{Title:title,Presentation:introduce,Sender:&c.user,Status:&status}
	err:=newSubject.Add()
	c.Data["json"]=err
	c.ServeJSON(true)

}

// 显示题目列表

func (c *TeacherController) SubjectList()  {

}



