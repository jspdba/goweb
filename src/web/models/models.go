package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)
//收藏的链接
type Link struct {
	Id int			`orm:"auto"`
	Url string		`orm:"unique;type(text)"`
	Title string		`orm:"null;type(text)"`
	Show bool		`orm:"default(1)"` //显示?1:0
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

func init() {
	// 需要在init中注册定义的model
	//maxIdle := 30
	//maxConn := 30
	orm.RegisterModel(new(Link),new(Tag))
}