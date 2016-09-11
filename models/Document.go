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

func (d *Document) Add() error{
	o := orm.NewOrm()

	_, err := o.Insert(d)
	return err

}

func (d *Document) Get() bool{
	o := orm.NewOrm()

	err := o.Read(d)
	return err == nil
}