package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.LoginRedirect()

	c.Data["Title"] = "登陆"
	c.TplName = "login.tpl"

}

func (c *LoginController) Post() {
	c.LoginRedirect()

	user_id := c.GetString("user_id")
	password := c.GetString("password")

	u := models.User{UserId: user_id, Password: password}

	switch u.Login() {
	case models.ERROR_USER_NOT_EXIST:
		c.ToError("用户名不存在","LoginController.Get")
	case models.ERROR_PASSWORD_ERROR:
		c.ToError("密码错误","LoginController.Get")
	case models.SUCCESS:
		c.SetSession("user_id", u.UserId)
		c.ToIndex()
	}

}

func (c *LoginController) Logout() {
	c.DelSession("user_id")
	c.ToIndex()

}

// 如果登陆则重定向到主页
func (c *LoginController) LoginRedirect() {
	if c.isLogin {
		c.ToIndex()
	}
}