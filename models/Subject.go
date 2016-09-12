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
	PaperGrade	float32
	DefenceGrade	float32
	FinalGrade 	float32 		`orm:"-"`
	Comment 	string
	//Time		time.Time
}

func init() {
	orm.RegisterModel(new(Subject))
}

func GetSubject(id int)(s *Subject , err error){
	s =&Subject{Id:id}
	err = s.Get()
	return s,err
}

func (s *Subject) Get() error {
	o := orm.NewOrm();
	err :=o.QueryTable("subject").Filter("Id",s.Id).RelatedSel().One(s)
	s.Count()
	return err
}

func (s *Subject) Count() {
	if s.PaperGrade == 0 {
		return
	}

	if s.DefenceGrade == 0 {
		return
	}

	s.FinalGrade = (s.PaperGrade + s.DefenceGrade) /2
}



func (c *Subject) SetStatus(status int) {
	o := orm.NewOrm();
	c.Status.Id = status
	o.Update(c,"Status")
}