package test

import (
	"testing"
	"github.com/sluu99/uuid"
	"crypto/md5"
	"web/models"
	"encoding/hex"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)


func Test_insert_user(t *testing.T) {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8", 30)
	orm.RunSyncdb("default", false, true)

	var token = uuid.Rand().Hex()
	h := md5.New()
	username := "jspdba"
	password := "wuchaofei1"
	h.Write([]byte(password))
	user := models.User{Username: username, Password: hex.EncodeToString(h.Sum(nil)), Avatar: "/static/imgs/avatar.png", Token: token}
	models.SaveUser(&user)
}
