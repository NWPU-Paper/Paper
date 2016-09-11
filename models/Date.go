package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Date struct {
	Id    		int
	Date  		time.Time
}

//获取日期信息 1为开始 2为结束 超过这个时间即不允许再布置题目了
//date := models.SetDate(2,time.Now())
func GetDate(date_type int) *Date{
	o := orm.NewOrm()
	date := &Date{Id:date_type}
	o.Read(date)
	return date
}

//管理员可以设置日期
//models.SetDate(2,time.Now())

func SetDate(date_type int,date time.Time) {
	o := orm.NewOrm()
	d := Date{Id:date_type,Date:date}
	o.Update(&d)
}

func init() {
	orm.RegisterModel(new(Date))
}
