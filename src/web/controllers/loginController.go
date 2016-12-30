package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) URLMapping() {
	this.Mapping("/user/list)", this.List)
}

 //@router /user/list
func (this *LoginController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.Page(page.PageNo,page.PageSize)
	this.TplName = "user/list.tpl"
}

