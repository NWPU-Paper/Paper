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
		err := c.user.GetUser()
		if err == nil {
			c.isLogin = true
		}
	}

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}

}

func (c *BaseController) ToError(msg string,ufos string)  {
	c.Ctx.WriteString("<h1>>(+_+)?<br>"+ msg +"</h1><script>var i = 3;function out(){if(i>1){i--;}else{location.href=\""+ beego.URLFor(ufos) + "\";}document.getElementById('aa').innerHTML=i+\"秒后跳转\";}setInterval('out()',1000);</script><p><div id=\"aa\">3秒后跳转</div></p>")
	c.StopRun()
}

func (c *BaseController) ToIndex() {
	c.RedirectTo("IndexController.Get")
}

func (c *BaseController) RedirectTo(endpoint string, values ...interface{}) {
	c.Redirect(c.URLFor(endpoint,values...), 302)
}

func (c *BaseController) NotFound() {
	c.ToError("没有找到这个文件","AdminController.Get")
}

func (c *BaseController) R()  {
	c.RedirectTo("SubjectController.Get",":id",c.GetString("id"))

}