package models

import (
	"github.com/astaxie/beego/orm"
)

//用户类型
const TYPE_USER_STUDENT = 1
const TYPE_USER_TEACHER = 2
const TYPE_USER_LEADER = 3
const TYPE_USER_SECRETARY = 4
const TYPE_USER_ADMIN = 5

//登陆错误信息
const ERROR_USER_NOT_EXIST = -1
const ERROR_PASSWORD_ERROR = -2
const SUCCESS = 0

type User struct {
	UserId   string		`orm:"pk"`
	Name     string
	Password string
	Email    string
	Gender   int
	Phone    int
	Major    *Major 	`orm:"rel(fk)"`
	Type     int
}

func init() {
	orm.RegisterModel(new(User))
}

func (c *User) Login() int {
	o := orm.NewOrm()
	data := *c
	err := o.Read(c)
	if err == orm.ErrNoRows {
		//用户不存在
		return ERROR_USER_NOT_EXIST
	} else if c.Password != data.Password {
		//密码不正确
		return ERROR_PASSWORD_ERROR
	} else {
		return SUCCESS
	}
}

func (c *User) GetUser() bool{
	o := orm.NewOrm()
	err := o.Read(c)
	if err == orm.ErrNoRows {
		//用户不存在
		return false
	}
	return true

}
