package controllers

import "github.com/astaxie/beego/orm"

type ProfileController struct {
	AdminController
}

func (c *ProfileController) LoginPrepare() {
	c.Data["Title"] = "个人信息"
	c.Layout = "base/base2.tpl"
	c.LayoutSections["Script"] = "base/script.tpl"
	c.TplName = "profile.tpl"
	c.Data["pageKey"] = "profile";

}


func (c *ProfileController) Get() {
	c.Data["Title"] = "个人信息"
	c.Layout = "base/base2.tpl"
	c.LayoutSections["Script"] = "base/script.tpl"
	c.TplName = "profile.tpl"
}

func (c *ProfileController) Post(){
	c.user.Email = c.GetString("email")
	c.user.Phone = c.GetString("phone")
	o :=orm.NewOrm()
	o.Update(&c.user)
	c.RedirectTo("ProfileController.Get")
}