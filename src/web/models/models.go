package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strings"
)
//收藏的链接
type Link struct {
	Id int			`orm:"auto"`
	Url string		`orm:"unique;type(text)"`
	Title string		`orm:"null;type(text)"`
	Show bool		`orm:"default(true)"` //显示?1:0
	Description string	`orm:"null;type(text)"`
	CreateDate time.Time	`orm:"auto_now_add;type(datetime)"`
	ModifyDate time.Time	`orm:"auto_now;type(datetime)"`
	Tags []*Tag 		`orm:"null;rel(m2m);on_delete(do_nothing)"`
}

//标签
type Tag struct {
	Id int			`orm:"auto"`
	Name string		`orm:"unique;type(text)"`
	CreateDate time.Time	`orm:"auto_now_add;type(datetime)"`
	Description string 	`orm:"null;type(text)"`
	Link []*Link 		`orm:"null;reverse(many);on_delete(do_nothing)"`
}

func TagReadOrCreate(name string) (tag *Tag){
	if name!=""{
		o := orm.NewOrm()
		tag=new(Tag)
		tag.Name=name
		if _, id, err := o.ReadOrCreate(tag, "Name"); err == nil {
			tag.Id=int(id)
		}else{
			beego.Error(err)
		}
	}
	return tag
}
func LinkReadOrCreate(link *Link,tags string) {
	o := orm.NewOrm()

	//save tags to db
	if tags!=""{
		tagArr:=strings.Split(tags,",")
		link.Tags=make([]*Tag,len(tagArr))
		for index,name := range tagArr{
			tag:=TagReadOrCreate(name)
			if tag!=nil{
				link.Tags[index]=tag
			}
		}
	}

	// 三个返回参数依次为：是否新创建的，对象Id值，错误
	if _, _, err := o.ReadOrCreate(link, "Url"); err == nil {
		if len(link.Tags)>0{
			beego.Info("Link.Tags size =",len(link.Tags))
			m2m := o.QueryM2M(link, "Tags")
			if _,err := m2m.Add(link.Tags); err!=nil{
				beego.Error(err)
			}
		}
	}else{
		beego.Error(err)
	}
}

func init() {
	// 需要在init中注册定义的model
	//maxIdle := 30
	//maxConn := 30
	orm.RegisterModel(new(Link),new(Tag))
}