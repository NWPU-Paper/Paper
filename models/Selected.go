package models

import "github.com/astaxie/beego/orm"

type Selected struct {
	Id int
}

func init()  {
	orm.RegisterModel(new(Selected))
}