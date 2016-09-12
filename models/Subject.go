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
	PaperGrade	float64
	DefenceGrade	float64
	FinalGrade 	float64 		`orm:"-"`
	Comment 	string
	//Time		time.Time
	Users	[]*User			`orm:"rel(m2m)"`
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
	//o.Read(s)
	err :=o.QueryTable("subject").Filter("Id",s.Id).RelatedSel().One(s)
	//fmt.Println(s.SelectUser)
	//if err == nil {
	//	 o.LoadRelated(s,"SelectUser")
	//	if err2 !=nil {
	//		fmt.Println(err2)
	//		fmt.Println(num)
	//	}
	//}
	//s.Users = s.GetSelectUser()
	//_,err := o.LoadRelated(s,"Users")
	s.Count()
	return err
}

func (s *Subject) GetSelectUser() (users []*User) {
	o := orm.NewOrm();
	o.QueryTable("user").Filter("Subjects__Subject__Id", s.Id).All(&users);
	return users
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

//func (s *Subject) GetSelect() (num int, err error){
//	o := orm.NewOrm()
//	num, err = o.Raw("SELECT * FROM `user` WHERE `user`.user_id in (SELECT user_id FROM subject_users WHERE  subject_users.subject_id = ?)",  s.Id ).ValuesList(s.SelectList)
//	return num, err
//}



func (c *Subject) SetStatus(status int) {
	o := orm.NewOrm();
	c.Status.Id = status
	o.Update(c,"Status")
}

func (c *Subject) SetGrade(t ,grade float64){
	switch t {
	case 0:
		//Paper Grade
		c.PaperGrade = grade
		break
	case 1:
		//Defence Grade
		c.DefenceGrade = grade
		break
	}
}

func (c *Subject) AddFile(t int ,d Document)  {
	o := orm.NewOrm()
	switch t {
	case TYPE_FILE_TASK:
		c.Task = &d
		break

	case TYPE_FILE_INTERIM:
		c.Interim = &d
		break

	case TYPE_FILE_PAPER:
		c.Paper = &d
		break

	case TYPE_FILE_TRANSLATE:
		c.Translate = &d
		break

	case TYPE_FILE_PRIMARY:
		c.Primary = &d
		break
	}
	o.Update(c)
}

func (d *Subject) Add() error {
	o := orm.NewOrm()
	d.Status.Id = 10
	_, err := o.Insert(d)
	return err

}