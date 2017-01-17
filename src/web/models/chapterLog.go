package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id int			`orm:"auto"`
	Tag string		`orm:"unique;type(text)"`
	Index string
	BookId string
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

func LogInsert(log *Log) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(log)
	return id
}

func FindLastLogByTagAndBookId(tag string,bookId string) (bool, Log) {
	o := orm.NewOrm()
	var log Log
	err := o.QueryTable(log).Filter("Tag", tag).Filter("BookId",bookId).OrderBy("-Index").Limit(1).One(&log)
	return err != orm.ErrNoRows, log
}

func LogInsertOrUpdate(log *Log) bool{
	o := orm.NewOrm()
	var obj Log
	err := o.QueryTable(obj).Filter("Tag", log.Tag).Filter("BookId",log.BookId).OrderBy("-Index").Limit(1).One(&obj)
	if err==nil{
		if obj.Index!=log.Index{
			obj.Index = log.Index
			if _,err:=o.Update(&obj);err==nil{
				return true
			}
		}
	}else{
		if _,err:=o.Insert(log); err==nil{
			return true
		}
	}
	return false;
}
func FindLogCount(tag string) int64{
	o := orm.NewOrm()
	var obj Log
	i,err := o.QueryTable(obj).Filter("tag", tag).Count()
	if err==nil{
		return i
	}
	return int64(0)
}

func init() {
	orm.RegisterModel(new(Log))
}