package controllers

import (
	"github.com/banixc/paper/models"
	"fmt"
)

type AuthController struct {
	BaseController
}

func (c *AuthController) Get() {
	c.Data["Title"] = "登陆"
	c.TplName = "login.html"
}

func (c *AuthController) Post() {
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		c.Out("解析错误")
	} else {
		status := u.Login()
		fmt.Println(status)
		if status > 0 {
			c.SetSession("user_id", status)
			c.Redirect(c.URLFor("ListController.Get"),301)
		} else {
			switch status {
			case -1:
				c.Out("用户名不存在")
				break
			case -2:
				c.Out("密码错误")
				break
			default:
				c.Out("未知错误")
				break
			}
		}
	}
}

func (c *AuthController) Logout() {
	c.SetSession("user_id", nil)
	c.Out("登出成功")
}



