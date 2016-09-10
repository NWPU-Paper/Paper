package models

import "github.com/astaxie/beego/orm"

type Major struct {
	Id    		int
	MajorName  	string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Major))
}