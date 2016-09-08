package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Out(message string) {
	c.Ctx.WriteString(message)
}

