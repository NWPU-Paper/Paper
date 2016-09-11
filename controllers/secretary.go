package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
)

type SecretaryController struct {
	AdminController
}

func (c *SecretaryController) LoginPrepare() {

	if c.user.Type != models.TYPE_USER_SECRETARY {
		c.Abort("403")
	}
	c.Data["Title"] = "答辩秘书"
}