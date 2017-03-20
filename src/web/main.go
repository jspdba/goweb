package main

import (
	_ "web/routers"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego"
	"web/service"
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
	localUrl:=beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8"
	remoteUrl:=beego.AppConfig.String("remote.mysqluser") + ":" + beego.AppConfig.String("remote.mysqlpass") + "@tcp("+beego.AppConfig.String("remote.host")+":"+beego.AppConfig.String("remote.port")+")/beego?charset=utf8"

	orm.RegisterDataBase("default", "mysql", localUrl, 30)
	orm.RegisterDataBase("remote", "mysql", remoteUrl, 1)
	orm.RunSyncdb("default", false, true)
	//切换数据库
	//o1 := orm.NewOrm()
	//o1.Using("remote")
}

func main() {
	orm.Debug = false

	//更新数据库job为初始状态
	service.ResetJob()

	/*beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"https:/*//*.foo.com"},
		AllowMethods: []string{"PUT", "PATCH"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))*/
	//http://stackoverflow.com/questions/28216342/how-to-set-acces-control-allow-origin-in-beego-framework
	/*beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET","PUT", "PATCH"},
		AllowHeaders: []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
	}))*/
	beego.Run()
	//orm.RunCommand()
}