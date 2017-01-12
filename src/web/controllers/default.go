package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["title"] = "私人空间"
	c.TplName = "index.tpl"
}
