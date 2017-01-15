package main

import (
	_ "web/routers"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego"
)

func init() {
	/*// 数据库别名
	name := "default"

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:wuchaofei@/beego?charset=utf8")*/

	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8", 30)
	orm.RunSyncdb("default", false, true)
}

func main() {
	//orm.Debug = true
	beego.Run()
	//orm.RunCommand()
}