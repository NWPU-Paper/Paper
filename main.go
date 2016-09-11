package main

import (
	_ "github.com/NWPU-Paper/Paper/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "paper:paper@tcp(mysql.banixc.com:3306)/paper?charset=utf8mb4,utf8")
	orm.RunSyncdb("default", false, true)

	//开启Debug
	orm.Debug = true
}

func main() {
	beego.AddFuncMap("Now",time.Now)

	beego.Run()
}
