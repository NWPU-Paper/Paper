package models

import (
	"github.com/astaxie/beego/orm"
)

//用户类型
const TYPE_USER_STUDENT = 1
const TYPE_USER_TEACHER = 2
const TYPE_USER_LEADER = 3
const TYPE_USER_SECRETARY = 4
const TYPE_USER_ADMIN = 5

//登陆错误信息
const ERROR_USER_NOT_EXIST = -1
const ERROR_PASSWORD_ERROR = -2
const SUCCESS = 0

type User struct {
	UserId   string		`orm:"pk"`
	Name     string
	Password string
	Email    string
	Gender   int
	Phone    string
	Major    *Major 	`orm:"rel(fk);null"`
	Type     int
	MaxStudent int
	NowStudent int
}

func init() {
	orm.RegisterModel(new(User))
}

// 用户登陆
func (c *User) Login() int {
	data := *c
	err := c.GetUser()
	if err == orm.ErrNoRows {
		//用户不存在
		return ERROR_USER_NOT_EXIST
	} else if c.Password != data.Password {
		//密码不正确
		return ERROR_PASSWORD_ERROR
	} else {
		return SUCCESS
	}
}

// 读取用户
func (c *User) GetUser() error{
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("UserId",c.UserId).RelatedSel().One(c)
	return err
}

//老师::获得自己发布的课题列表

func GetAllSubject(teacher_name string)(subjects []Subject) {
	o := orm.NewOrm()
	if teacher_name == "" {
		o.QueryTable("subject").RelatedSel().OrderBy("-id").All(&subjects)
	} else {
		o.QueryTable("subject").Filter("Sender__Name__contains",teacher_name).OrderBy("-id").RelatedSel().All(&subjects)
	}
	return subjects
}

func (c *User) GetSendSubject()(subjects []Subject) {
	o := orm.NewOrm()
	o.QueryTable("subject").Filter("sender_id",c.UserId).OrderBy("-id").RelatedSel().All(&subjects)
	return subjects
}

//学生::获得当前能选择的课题
func (c *User) GetSubjectList(teacher_name string)(subjects []Subject) {
	o := orm.NewOrm()

	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("subject.id").
		From("subject").
		InnerJoin("user sender").On("sender.user_id = subject.sender_id")

	if teacher_name != "" {
		qb.Where("sender.name LIKE '%"+ teacher_name +"%' AND sender.major_id = ?")
	} else {
		qb.Where("sender.major_id = ?")
	}


	sql := qb.String()
	o.Raw(sql,c.Major.Id).QueryRows(&subjects)
	return GetSelectListByList(subjects)
}

//学生::获得当前选择的课题
func (c *User) GetSelectSubject()(subjects []Subject) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("id").
		From("subject").
		//InnerJoin("selected").On("subject.id = selected.subject_id").
		//InnerJoin("status").On("status.id = subject.status_id").
		Where("student_id = ?")
	sql := qb.String()

	o.Raw(sql,c.UserId).QueryRows(&subjects)

	return GetSelectListByList(subjects)
}


func GetSelectListByList(s []Subject)(subjects []Subject){
	o := orm.NewOrm()
	var l = []int{}
	for _, r := range s {
		l = append(l,r.Id)
	}
	if len(l) !=0 {
		o.QueryTable("subject").Filter("Id__in", l).RelatedSel().OrderBy("-id").All(&subjects)
	}
	return subjects

}

////学生::获得当前锁定的课题
//func (c *User) GetLockedSubject() (subject Subject,err error) {
//	o := orm.NewOrm()
//	subject = Subject{Student:&User{UserId:c.UserId}}
//	err = o.Read(&subject,"student_id")
//	return subject,err
//}

func (c *User) SelectSubject(id int) error {
	o := orm.NewOrm()

	sql := "INSERT INTO `paper`.`selected` (`user_id`, `subject_id`) VALUES (?, ?);"

	_, err := o.Raw(sql,c.UserId,id).Exec()
	return err
}

