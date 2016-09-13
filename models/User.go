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
	Major    *Major 	`orm:"rel(fk)"`
	Type     int
	MaxStudent int
	NowStudent int
}

func init() {
	orm.RegisterModel(new(User))
}

// 用户登陆
func (c *User) Login() int {
	o := orm.NewOrm()
	data := *c
	err := o.Read(c)
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

// 读取用户 没啥用 打算删了
func (c *User) GetUser() bool{
	o := orm.NewOrm()
	err := o.Read(c)
	if err == orm.ErrNoRows {
		//用户不存在
		return false
	}
	return true
}

//老师::获得自己发布的课题列表
func (c *User) GetSendSubject()(subjects []Subject) {
	o := orm.NewOrm()
	o.QueryTable("subject").Filter("sender_id",c.UserId).All(&subjects)
	return subjects
}

//学生::获得当前选择的课题
func (c *User) GetSelectSubject()(subjects []Subject) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("subject.*").
		From("subject").
		InnerJoin("selected").On("subject.id = selected.subject_id").
		Where("user_id = ?")
	sql := qb.String()

	o.Raw(sql,c.UserId).QueryRows(&subjects)
	return subjects
}

//学生::获得当前锁定的课题
func (c *User) GetLockedSubject() (subject Subject,err error) {
	o := orm.NewOrm()
	subject = Subject{Student:&User{UserId:c.UserId}}
	err = o.Read(&subject,"student_id")
	return subject,err
}

func (c *User) SelectSubject(id int) error {
	o := orm.NewOrm()

	sql := "INSERT INTO `paper`.`selected` (`user_id`, `subject_id`) VALUES (?, ?);"

	_, err := o.Raw(sql,c.UserId,id).Exec()
	return err
}

