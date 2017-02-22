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
		Ip:"111.124.7.14",
		Port:"9999",
	},3*time.Second)
	beego.Info(valid)
}
