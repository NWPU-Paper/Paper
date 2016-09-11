package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/NWPU-Paper/Paper/models"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {


	o := orm.NewOrm();
	u := models.User{UserId:"3"}
	o.QueryTable("user").Filter("UserId",1).RelatedSel().One(&u)
	c.Data["json"] = u
	c.ServeJSON()
}
