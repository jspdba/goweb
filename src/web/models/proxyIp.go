package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"web/utils"
)

type ProxyIp struct {
	Id int			`orm:"auto"`
	Ip string
	Port string		`orm:"default(80)"`
	Typ string		`orm:"default(http)"`
	Area string		`orm:"null;type(text)"`
	ConnSecond int		`orm:"null;default(0)"`
	Description string	`orm:"null;type(text)"`
	Useful bool		`orm:"defalt(true)"`
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

func ProxyIpSave(proxyIp *ProxyIp) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(proxyIp)
	return id
}
func ProxyIpSaveMulti(c [] *ProxyIp) int64 {
	o := orm.NewOrm()
	successNums, err:= o.InsertMulti(len(c),c)
	if err!=nil{
		beego.Error(err)
	}
	return successNums
}
func ProxyIpSaveOrUpdate(proxyIp *ProxyIp) int64{
	o := orm.NewOrm()
	proxyIpOld:=*proxyIp
	if created, id, err := o.ReadOrCreate(proxyIp, "Ip","Port"); err == nil {
		if created {
			return id
		} else {
			proxyIp.Ip = proxyIpOld.Ip
			proxyIp.Port = proxyIpOld.Port
			proxyIp.Typ = proxyIpOld.Typ
			proxyIp.Area = proxyIpOld.Area
			proxyIp.ConnSecond = proxyIpOld.ConnSecond
			proxyIp.Description = proxyIpOld.Description
			proxyIp.CreateDate = proxyIpOld.CreateDate
			if id, err := o.Update(proxyIp); err==nil{
				return id
			}else{
				beego.Error(err)
			}
		}
	}
	return int64(-1)
}
func ProxyIpDelete(proxyIp *ProxyIp) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.Delete(&proxyIp); err == nil {
		if num>0{
			result=true
		}
	}
	return result
}
func ProxyIdDeleteById(id int) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.QueryTable("proxy_id").Filter("id__eq", id).Delete(); err == nil {
		if num>0{
			result=true
		}
	}
	return result
}
func FindProxyIpById(id int) (bool, ProxyIp) {
	o := orm.NewOrm()
	var proxyIp ProxyIp
	err := o.QueryTable(proxyIp).Filter("Id", id).One(&proxyIp)
	return err != orm.ErrNoRows, proxyIp
}
func FindProxyIpByStrId(id string) (bool, ProxyIp) {
	o := orm.NewOrm()
	var proxyIp ProxyIp
	err := o.QueryTable(proxyIp).Filter("Id", id).One(&proxyIp)
	return err != orm.ErrNoRows, proxyIp
	return false,proxyIp
}

func ProxyIpPage(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj ProxyIp
	var list []ProxyIp
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
func init() {
	orm.RegisterModel(new(ProxyIp))
}