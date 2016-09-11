package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "paper"
	c.Data["Email"] = "banixc@gmail.com"
	c.TplName = "index.tpl"

}

func (c *MainController) Clear() {
	c.DelSession("user_id")
	c.TplName = "index.tpl"
}
