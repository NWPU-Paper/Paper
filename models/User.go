package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type User struct {
	Id    int		`form:"-"`
	Name  string		`form:"-"`
	Username string		`form:"username"`
	Password string		`form:"password"`
	Email string		`form:"-"`
}

//func init() {
//	// 需要在init中注册定义的model
//}

func (c *User) Login() int {
	o := orm.NewOrm()
	data := User{Username:c.Username}
	err := o.Read(&data, "Username")
	if err == orm.ErrNoRows {
		//用户名不存在
		return -1
	} else if c.Password != data.Password {
		fmt.Println(c.Password)
		fmt.Println(data.Password)
		//密码不正确
		return -2
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
