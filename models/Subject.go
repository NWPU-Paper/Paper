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
	Student 	*User		`orm:"rel(one);null"`
	PaperGrade	float64
	DefenceGrade	float64
	FinalGrade 	float64 		`orm:"-"`
	Comment 	string
	//Time		time.Time
	Users		[]User		`orm:"-"`
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
	s.Users = s.GetSelectUser()
	s.Count()
	return err
}


func (s *Subject) GetSelectUser() (users []User) {

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("user.*").
		From("user").
		InnerJoin("selected").On("user.user_id = selected.user_id").
		Where("subject_id = ?")
	sql := qb.String()

	o := orm.NewOrm();
	o.Raw(sql, s.Id).QueryRows(&users)
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

//老师/负责人::锁定当前学生
func (s *Subject) Lock(user_id string) error {
	o := orm.NewOrm()

	s.Student = &User{UserId:user_id}
	s.Status = &Status{Id:STATUS_SELECTED_LOCKED}
	_,err := o.Update(s)
	if err != nil {
		return err
	}
	//删除所有选定的人
	_,err = o.QueryTable("selected").Filter("user_id", user_id).Delete()
	if (err != nil) {
		return err
	}
	_,err = o.QueryTable("selected").Filter("subject_id",s.Id).Delete()
	return err
}