package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"web/utils"
)

type Job struct {
	Id int			`orm:"auto"`
	Name string		`orm:"unique;"`
	Cron string		`orm:"null;type(text)"`
	Content string		`orm:"null;type(text)"`
	State int		`orm:"default(0)"`
	BookId int		`orm:"null"`
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

func FindJobById(id string) (bool ,*Job) {
	o := orm.NewOrm()
	var job Job
	i,er:=strconv.Atoi(id)
	if(er!=nil){
		return false,&job
	}
	err := o.QueryTable(job).Filter("Id", i).One(&job)
	return err != orm.ErrNoRows, &job
}

func JobInsert(job *Job) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(job)
	return id
}

func JobUpdate(job *Job) int64{
	o := orm.NewOrm()
	if id, err := o.Update(job); err==nil{
		return id
	}else{
		beego.Error(err)
	}
	return int64(0)
}

func JobDelete(job *Job) (bool,int64) {
	o := orm.NewOrm()
	if num, err := o.Delete(&job); err == nil {
		return true,num
	}
	return false,0
}
func JobDeleteById(id string) (bool,int64) {
	o := orm.NewOrm()
	if id!=""{
		if i,err:=strconv.Atoi(id);err==nil{
			job:=Job{Id:i}
			if num, err := o.Delete(&job); err == nil {
				return true,num
			}
		}
	}
	return false,0
}

func JobSaveOrUpdate(job *Job) int64 {
	o := orm.NewOrm()
	jobOld:=*job
	if created, id, err := o.ReadOrCreate(job, "Id"); err == nil {
		if created {
			return id
		} else {
			job.Name = jobOld.Name
			job.Cron = jobOld.Cron
			job.Content = jobOld.Content
			job.State = jobOld.State
			job.BookId = jobOld.BookId
			if id, err := o.Update(job); err==nil{
				return id
			}else{
				beego.Error(err)
			}
		}
	}
	return int64(-1)
}

func JobPage(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj Job
	var list []Job
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func init() {
	orm.RegisterModel(new(Job))
}