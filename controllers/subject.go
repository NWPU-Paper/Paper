package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
	"strconv"
)

type SubjectController struct {
	AdminController
}

func (c *SubjectController) LoginPrepare() {
	c.Data["Title"] = "课题"
	c.Layout = "base/base2.tpl"
}

func (c *SubjectController) Get() {

}

func (c *SubjectController) Detail(){
	id ,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	s,err:= models.GetSubject(id)
	if err== nil {
		c.TplName="subject/body.tpl"
		c.Data["Subject"] = s;
		c.Data["json"] = s;
		//c.ServeJSON();
	} else {
		c.NotFound()
	}
}

