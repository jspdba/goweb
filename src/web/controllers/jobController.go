package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type JobController struct {
	beego.Controller
}

func (this *JobController) URLMapping() {
	this.Mapping("/job/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/job/save", this.SaveOrUpdate)
	this.Mapping("/job/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/job/list", this.List)
}
// @router /job/edit/:id([0-9]{0,}) [get]
func (this *JobController) Edit() {
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		_,job:=models.FindJobById(id)
		this.Data["entity"] = job
	}
	this.TplName = "job/edit.tpl"
}

// @router /job/save [post]
func (this *JobController) SaveOrUpdate() {
	job := models.Job{}
	if err := this.ParseForm(&job); err != nil {
		beego.Error(err)
	}
	models.JobSaveOrUpdate(&job)
	this.Redirect("/job/list", 302)
}

// @router /job/delete/:id([0-9]+)
func (this *JobController) Delete(){
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		models.JobDeleteById(id)
	}
}

// @router /job/list
func (this *JobController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.JobPage(page.PageNo,page.PageSize)
	this.TplName = "job/list.tpl"
}