package test

import (
	"testing"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"web/service"
)

func TestImportTableLink(t *testing.T) {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8", 30)
	var remote=beego.AppConfig.String("remote.mysqluser") + ":" + beego.AppConfig.String("remote.mysqlpass") + "@tcp("+beego.AppConfig.String("remote.host")+":"+beego.AppConfig.String("remote.port")+")/beego?charset=utf8"
	orm.RegisterDataBase("remote", "mysql", remote)
	orm.RunSyncdb("default", false, true)
	service.ImportRemoteLinkTable()
}
func TestImportRemoteBookTable(t *testing.T) {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8", 30)
	var remote=beego.AppConfig.String("remote.mysqluser") + ":" + beego.AppConfig.String("remote.mysqlpass") + "@tcp("+beego.AppConfig.String("remote.host")+":"+beego.AppConfig.String("remote.port")+")/beego?charset=utf8"
	orm.RegisterDataBase("remote", "mysql", remote)
	orm.RunSyncdb("default", false, true)
	service.ImportRemoteBookTable()
}
