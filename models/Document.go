package models

import "github.com/astaxie/beego/orm"

type Document struct {
	Id 	int
	Name	string
	Path 	string
}

func init() {
	orm.RegisterModel(new(Document))
}