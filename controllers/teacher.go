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
	c.TplName = "teacher/index.tpl"


}

// 发布新题目
func (c *TeacherController) Post() {

}

// 显示题目列表

func (c *TeacherController) SubjectList()  {

}

// 显示题目详细

func (c *TeacherController) Subject()  {

}

