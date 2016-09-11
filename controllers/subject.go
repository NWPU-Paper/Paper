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

}

func (c *SubjectController) Get() {

}

func (c *SubjectController) Detail(){
	id ,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	s := models.GetSubject(id)
	if s!= nil {
		c.Data["json"] = s
		c.ServeJSON()
	} else {
		c.NotFound()
	}
}

