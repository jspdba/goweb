package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"web/utils"
)

//用户entity
type User struct {
	Id int				`orm:"auto"`
	Username string
	Password string
	Mobile string			`orm:"null;"`
	Status int			`orm:"default(0)"`	//用户状态(0=正常，1=停用)
	Mail string			`orm:"null;"`
	CreateDate  time.Time 		`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 		`orm:"auto_now;type(datetime)"`
}

//https://github.com/lyonlai/bootstrap-paginator
func Page(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj User
	var list []User
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
func init() {
	// 需要在init中注册定义的model
	//maxIdle := 30
	//maxConn := 30
	orm.RegisterModel(new(User))
}