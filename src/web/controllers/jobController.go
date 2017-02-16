package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
	"web/job"
)

type JobController struct {
	beego.Controller
}

func (this *JobController) URLMapping() {
	this.Mapping("/job/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/job/save", this.SaveOrUpdate)
	this.Mapping("/job/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/job/list", this.List)
	this.Mapping("/job/start/:id([0-9]+)", this.Start)
	this.Mapping("/job/pause/:id([0-9]+)", this.Pause)
	this.Mapping("/job/run/:id([0-9]+)", this.Run)
}
// @router /job/edit/:id([0-9]{0,}) [get]
func (this *JobController) Edit() {
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		_,job:=models.FindJobById(id)
		this.Data["entity"] = job
	}

	page:=utils.Page{PageNo:1,PageSize:20}
	this.Data["page"] = models.BookPage(page.PageNo,page.PageSize)
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

// 启动任务
//@router /job/start/:id([0-9]+)
func (this *JobController) Start() {
	id:=this.Ctx.Input.Param(":id")
	ok,task := models.FindJobById(id)
	if !ok {
		return
	}
	okk,book:=models.FindBookByStrId(id)
	if !okk{
		return
	}

	j, err := job.NewJobFromDb(task,&book)
	if err != nil {
		beego.Error(err)
	}

	if job.AddJob(task.Cron, j) {
		task.State = 1
		models.JobSaveOrUpdate(task)
	}

	refer := this.Ctx.Request.Referer()
	if refer == "" {
		refer = beego.URLFor("JobController.List")
	}
	this.Redirect(refer,302)
}

// 暂停任务
// @router /job/pause/:id([0-9]+)
func (this *JobController) Pause() {
	id:=this.Ctx.Input.Param(":id")

	ok,task := models.FindJobById(id)
	if !ok {
		return
	}

	job.RemoveJob(id)
	task.State = 0
	models.JobSaveOrUpdate(task)

	refer := this.Ctx.Request.Referer()
	if refer == "" {
		refer = beego.URLFor("JobController.List")
	}
	this.Redirect(refer,302)
}

// 立即执行
// @router /job/run/:id([0-9]+)
func (this *JobController) Run() {
	id:=this.Ctx.Input.Param(":id")

	ok,task := models.FindJobById(id)
	if !ok {
		return
	}
	okk,book:=models.FindBookByStrId(id)
	if !okk{
		return
	}

	j, err := job.NewJobFromDb(task,&book)
	if err != nil {
		beego.Error(err)
	}
	go j.Run()

	this.Redirect(beego.URLFor("JobController.List"), 302)
}