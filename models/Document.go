package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const TYPE_FILE_TASK = 1
const TYPE_FILE_PRIMARY = 2
const TYPE_FILE_INTERIM = 3
const TYPE_FILE_PAPER = 4
const TYPE_FILE_TRANSLATE = 5

type Document struct {
	Id 	int
	Name	string
	Path 	string
	Time 	time.Time
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