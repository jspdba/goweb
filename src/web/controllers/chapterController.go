package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"strconv"
	"web/utils"
)

type ChapterController struct {
	beego.Controller
}


func (this *ChapterController) URLMapping() {
	this.Mapping("/chapter/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/chapter/detail/:id([0-9]+)", this.Detail)
	this.Mapping("/chapter/save", this.SaveOrUpdate)
	this.Mapping("/chapter/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/chapter/list/:id([0-9]+)", this.List)
}

// @router /chapter/edit/:id([0-9]{0,}) [get]
func (this *ChapterController) Edit() {
	id:=this.Ctx.Input.Param(":id")
	obj:= models.Chapter{}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,entity:=models.FindChapterById(i)
			beego.Info(entity)
			if ok{
				obj=entity
			}
		}
	}
	this.Data["entry"] = obj
	this.TplName = "chapter/edit.tpl"
}
// @router /chapter/detail/:id([0-9]{0,}) [get]
func (this *ChapterController) Detail() {
	id:=this.Ctx.Input.Param(":id")
	obj:= models.Chapter{}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,entity:=models.FindChapterById(i)
			beego.Info(entity)
			if ok{
				obj=entity
			}
		}
	}
	this.Data["entity"] = obj
	this.TplName = "chapter/detail.tpl"
}

// @router /chapter/save [post]
func (this *ChapterController) SaveOrUpdate() {
	entity := models.Chapter{}
	if err := this.ParseForm(&entity); err != nil {
		beego.Error(err)
	}
	models.ChapterSaveOrUpdate(&entity)
	this.Redirect("/chapter/list", 302)
}

// @router /chapter/delete/:id([0-9]+)
func (this *ChapterController) Delete(){
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		if i,err:=strconv.Atoi(id); err==nil{
			entity := models.Chapter{Id:i}
			models.ChapterDelete(&entity)
		}

	}
}

// @router /chapter/list/:id([0-9]{0,}) [get]
func (this *ChapterController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}

	id:=this.Ctx.Input.Param(":id")
	bookId:=-1
	if id!="" {
		if i, err := strconv.Atoi(id); err == nil {
			bookId=i
		}
	}
	this.Data["page"] = models.ChapterPage(page.PageNo,page.PageSize,bookId)
	this.TplName = "chapter/list.tpl"
}