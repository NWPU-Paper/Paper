package controllers

type ListController struct {
	AdminController
}


func (c *ListController) Get(){
	c.Out("<html>这是管理的后台页面 只有登陆成功后才能进入</br>您的姓名为：" + c.CurrentUser.Name  + "</br><a href=\"" + c.URLFor("AuthController.Logout") + "\">登出</a></html>")
}
