package controllers

import "github.com/NWPU-Paper/Paper/models"

type ListController struct {
	AdminController
}

func (c *ListController) LoginPrepare() {
	c.Data["Title"] = "课题列表"
	c.Layout = "base/base2.tpl"
	c.LayoutSections["Script"] = "subject_detail/script.tpl"
	c.TplName = "list.tpl"

}

func (c *ListController) Get() {
	// 获取各类参数
	SearchTeacher := c.GetString("subject_sender")

	c.Data["SearchTeacher"] = SearchTeacher



	var subjects = []models.Subject{}
	switch c.user.Type {
	case models.TYPE_USER_STUDENT:
		//学生只能看到当前的专业
		subjects = c.user.GetSubjectList(SearchTeacher)
		c.Data["SubjectList"] = subjects
		break
	case models.TYPE_USER_TEACHER:
		//老师只能看到自己布置的专业
		subjects = c.user.GetSendSubject()

		break
	case models.TYPE_USER_LEADER:
		//领导只能看到自己专业的专业
		subjects = c.user.GetSubjectList(SearchTeacher)

		break

	//默认的秘书和管理员能看到所有的东西
	case models.TYPE_USER_SECRETARY:
		subjects = models.GetAllSubject(SearchTeacher)
		break

	case models.TYPE_USER_ADMIN:
		subjects = models.GetAllSubject(SearchTeacher)

		break
	default:
		c.Abort("403")

	}

	c.Data["SubjectList"] = subjects


}
