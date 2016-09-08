package main

import (
	_ "github.com/banixc/paper/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "paper:paper@/paper?charset=utf8mb4,utf8")
	orm.RunSyncdb("default", false, true)

	//开启Debug
	orm.Debug = true
}

func main() {
	beego.Run()
}

