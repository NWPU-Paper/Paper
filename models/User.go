package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    int
	Name  string		`orm:"size(64)"`
	Username string		`orm:"size(64)"`
	Password string		`orm:"size(64)"`
	Type int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

func (c *User) Login() int {
	o := orm.NewOrm()
	data := User{Username:c.Username}
	err := o.Read(&data, "Username")
	if err == orm.ErrNoRows {
		//用户名不存在
		return -1
	} else if c.Password != data.Password {
		//密码不正确
		return -2
	} else if c.Type != data.Type {
		//用户类型不正确
		return -3
	} else {
		return data.Id
	}
}

func (c *User) GetUser() {
	o := orm.NewOrm()
	err := o.Read(c)
	if err == orm.ErrNoRows {
		//用户不存在
		c.Id = -1
	}

}
