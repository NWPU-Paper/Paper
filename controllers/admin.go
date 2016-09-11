package controllers

type LoginPreparer interface{
	LoginPrepare()
}


//所有需要登陆的都继承这个 未登陆则自动返回
type AdminController struct {
	BaseController
}

func (c *AdminController) NestPrepare() {
	if !c.isLogin {
		//未登录则返回401
		c.Abort("401")
	}

	if app, ok := c.AppController.(LoginPreparer); ok {
		app.LoginPrepare()
	}

	c.Data["CurrentUser"] = c.user

}

