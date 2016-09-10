package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) NestPrepare() {

	//对于不同的用户导向不同的页面

	if !c.isLogin {
		c.RedirectTo("LoginController.Get")
	}
	switch c.user.Type {
		case models.TYPE_USER_STUDENT:
			c.RedirectTo("StudentController.Get")
			break
		case models.TYPE_USER_TEACHER:
			c.RedirectTo("TeacherController.Get")
			break
		case models.TYPE_USER_LEADER:
			break
		case models.TYPE_USER_SECRETARY:
			break
		case models.TYPE_USER_ADMIN:
			break
	}
}

