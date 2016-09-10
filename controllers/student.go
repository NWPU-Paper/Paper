package controllers

import "github.com/NWPU-Paper/Paper/models"

type StudentController struct {
	AdminController
}

func (c *StudentController) LoginPrepare() {

	if c.user.Type != models.TYPE_USER_STUDENT {
		c.Abort("403")
	}
}

func (c *StudentController) Get() {
	c.Ctx.WriteString("这是学生界面")
}