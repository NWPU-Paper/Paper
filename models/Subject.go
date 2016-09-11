package models

import (
	"github.com/astaxie/beego/orm"
)

type Subject struct {
	Id 		int
	Title		string
	Presentation 	string
	Sender 		*User 		`orm:"rel(fk)"`
	Task 		*Document	`orm:"rel(fk);null"`
	Primary 	*Document	`orm:"rel(fk);null"`
	Interim 	*Document	`orm:"rel(fk);null"`
	Paper	 	*Document	`orm:"rel(fk);null"`
	Translate	*Document	`orm:"rel(fk);null"`
	Status		*Status		`orm:"rel(fk)"`
	Student 	*User		`orm:"rel(fk);null"`
	PaperGrade	int
	DefenceGrade	int
	Comment 	string
	//Time		time.Time
}

func init() {
	orm.RegisterModel(new(Subject))
}

func GetSubject(id int) *Subject{
	s :=&Subject{Id:id}
	s.Get()
	return s
}

func (s *Subject) Get() {
	o := orm.NewOrm();
	o.QueryTable("subject").Filter("Id",s.Id).RelatedSel().One(s)
}



