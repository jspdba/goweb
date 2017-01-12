package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type BookController struct {
	beego.Controller
}

func (this *BookController) URLMapping() {
	this.Mapping("/book/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/book/save", this.SaveOrUpdate)
	this.Mapping("/book/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/book/list", this.List)
}

// @router /book/edit/:id([0-9]{0,}) [get]
func (this *BookController) Edit() {
	this.TplName = "book/edit.tpl"
}

// @router /book/save [post]
func (this *BookController) SaveOrUpdate() {
	book := models.Book{}
	if err := this.ParseForm(&book); err != nil {
		beego.Error(err)
	}
	models.BookSaveOrUpdate(&book)
	this.Redirect("/book/list", 302)
}

// @router /book/delete/:id
func (this *BookController) Delete() {
}

// @router /book/list
func (this *BookController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.LinkPage(page.PageNo,page.PageSize)
	this.TplName = "book/list.tpl"
}