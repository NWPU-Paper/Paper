package models

import "github.com/astaxie/beego/orm"

//出题阶段：已发布/审核成功/未通过
const STATUS_START  = 10
const STATUS_START_SUCCESS  = 11
const STATUS_START_FAIL  = 12

//上传计划书/通过/未通过
const STATUS_TASK_UPLOAD  = 20
const STATUS_TASK_SUCCESS  = 21
const STATUS_TASK_FAIL  = 22

//上传开题报告/通过/未通过
const STATUS_PRIMARY_UPLOAD  = 30
const STATUS_PRIMARY_SUCCESS  = 31
const STATUS_PRIMARY_FAIL = 32

//上传中期报告/通过/未通过
const STATUS_INTERIM_UPLOAD  = 40
const STATUS_INTERIM_SUCCESS  = 41
const STATUS_INTERIM_FAIL = 42

//上传论文/通过/未通过
const STATUS_PAPER_UPLOAD  = 50
const STATUS_PAPER_SUCCESS  = 51
const STATUS_PAPER_FAIL = 52

//上传论文译文/通过/未通过
const STATUS_TRANSLATE_UPLOAD  = 60
const STATUS_TRANSLATE_SUCCESS  = 61
const STATUS_TRANSLATE_FAIL = 62

//已答辩
const STATUS_DEFENCE  = 70
const STATUS_DEFENCE_FAIL  = 71

type Status struct {
	Id 	int
	Name 	string
}


func init() {
	orm.RegisterModel(new(Status))
}