package controllers

import (
	"github.com/NWPU-Paper/Paper/models"
	"strconv"
	"time"
	"os"
	"fmt"
)

const ACTION_CHANGE_STATUS = 1;
const ACTION_SET_PAPER_GRADE = 2;
const ACTION_SET_DEFENCE_GRADE = 3;

const ACTION_UPLOAD_FILE = 4;

const ACTION_SELECT_SUBJECT = 5;

const ACTION_LOCKED_SUBJECT = 6;

const ACTION_EDIT_SUBJECT = 7;

type SubjectController struct {
	AdminController
}

func (c *SubjectController) LoginPrepare() {
	c.Data["Title"] = "课题"
	c.Layout = "base/base2.tpl"
	c.LayoutSections["Script"] = "subject_detail/script.tpl"
}

func (c *SubjectController) Get() {
	id ,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	s,err:= models.GetSubject(id)
	if err== nil {
		c.TplName="subject_detail/subject_detail.tpl"
		c.Data["Subject"] = s;
		c.Data["json"] = s;
		//c.ServeJSON();
	} else {
		c.NotFound()
	}
}

func (c *SubjectController) Post() {
	id ,_ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	s,_:= models.GetSubject(id)

	action, err := c.GetInt("operate")
	if err == nil {
		switch action {
		case ACTION_CHANGE_STATUS:
			status, _ := c.GetInt("status")
			s.SetStatus(status)
			break
		case ACTION_SET_PAPER_GRADE:
			grade,_ := c.GetFloat("paper_grade")
			s.SetGrade(0,grade)
			break
		case ACTION_SET_DEFENCE_GRADE:
			grade,_ := c.GetFloat("defence_grade")
			s.SetGrade(1,grade)
			break
		case ACTION_UPLOAD_FILE:
			//Save File
			f, h, _ := c.GetFile("file")
			defer f.Close()
			document := models.Document{Name:h.Filename,Path:"./static/upload/" + time.Now().Format("2006/01")}

			os.MkdirAll(document.Path, 0644)

			path := fmt.Sprintf("%s/%s", document.Path, document.Name)

			c.SaveToFile("file",path)

			document.Path = document.Path[1:len(document.Path)] + "/"

			document.Add()

			t,_ := c.GetInt("file_type")

			s.AddFile(t,document)
			break

		case ACTION_SELECT_SUBJECT:
			models.SelectSubject(c.user.UserId,id)

		case ACTION_LOCKED_SUBJECT:
			user_id := c.GetString("locked_user_id")
			s.Lock(user_id)
		case ACTION_EDIT_SUBJECT:
			s.Title = c.GetString("subject_title")
			s.Presentation = c.GetString("subject_presentation")
			s.Save()
		default:
			c.Abort("503")


		}
		c.RedirectTo("SubjectController.Get" , ":id",id)

	} else {
		c.Abort("503")
	}


}