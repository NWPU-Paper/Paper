package controllers

import (
	"github.com/astaxie/beego"
	"github.com/NWPU-Paper/Paper/models"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
	user models.User
	isLogin bool
}

func (c *BaseController) Prepare() {

	id := c.GetSession("user_id")

	if id != nil {
		c.user = models.User{UserId:id.(string)}
		if c.user.GetUser() {
			c.isLogin = true
		}
	}

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

}

func (c *BaseController) ToIndex() {
	c.RedirectTo("IndexController.Get")
}

func (c *BaseController) RedirectTo(endpoint string, values ...interface{}) {
	c.Redirect(c.URLFor(endpoint,values), 302)
}

func (c *BaseController) Out(message string) {
	c.Ctx.WriteString(message)
}

func (c *BaseController) NotFound() {
	c.Ctx.WriteString("NotFound")
}