package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
	"strconv"
)

type ProxyIpController struct {
	beego.Controller
}

func (this *ProxyIpController) URLMapping() {
	this.Mapping("/proxyIp/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/proxyIp/save", this.Save)
	this.Mapping("/proxyIp/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/proxyIp/list", this.List)
}

// @router /proxyIp/edit/:id([0-9]{0,}) [get]
func (this *ProxyIpController) Edit() {
	this.TplName = "proxyIp/edit.tpl"
}

// @router /proxyIp/save [post]
func (this *ProxyIpController) Save() {
	proxyIp := models.ProxyIp{}
	if err := this.ParseForm(&proxyIp); err != nil {
		beego.Error(err)
	}
	models.ProxyIpSaveOrUpdate(&proxyIp)
	this.Redirect("/proxyIp/list", 302)
}

// @router /proxyIp/delete/:id
func (this *ProxyIpController) Delete() {
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		if i,er:=strconv.Atoi(id);er==nil{
			proxyIp := models.ProxyIp{Id:i}
			models.ProxyIpDelete(&proxyIp)
		}
	}

	this.Redirect("/proxyIp/list", 302)
}

// https://beego.me/docs/utils/page.md
// @router /proxyIp/list
func (this *ProxyIpController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.ProxyIpPage(page.PageNo,page.PageSize)
	this.TplName = "proxyIp/list.tpl"
}