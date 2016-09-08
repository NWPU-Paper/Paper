package controllers

import (
	"github.com/banixc/paper/models"

)

type AdminController struct {
	BaseController
	CurrentUser models.User
}

func (c *AdminController) Prepare()  {
	c.GetCurrentUser()
	if c.CurrentUser.Id == -1 {
		//未登录则返回403
		c.Abort("403")
	}
}

func (c *AdminController) GetCurrentUser() {
	id := c.GetSession("user_id")
	if id != nil {
		c.CurrentUser = models.User{Id:id.(int)}
		c.CurrentUser.GetUser()
	} else {
		c.CurrentUser = models.User{Id:-1}
	}
}



