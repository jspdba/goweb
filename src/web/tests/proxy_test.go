package test

import (
	"testing"
	"web/service"
	"web/models"
	"github.com/astaxie/beego"
	"time"
)

func  TestProxyIp(t *testing.T) {
	valid:=service.CheckProxyValid(&models.ProxyIp{
		Ip:"200.92.152.130",
		Port:"8080",
	},5*time.Second)
	beego.Info(valid)
}
