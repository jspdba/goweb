package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	"web/utils"
	"strconv"
)

//收藏的链接
type Link struct {
	Id          int       `orm:"auto"`
	Url         string    `orm:"unique;type(text)"`
	Title       string    `orm:"null;type(text)"`
	Show        bool      `orm:"default(true)"` //显示?1:0
	Description string    `orm:"null;type(text)"`
	CreateDate  time.Time `orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time `orm:"auto_now;type(datetime)"`
	Tags        []*Tag    `orm:"null;rel(m2m);on_delete(do_nothing)"`
}

//标签
type Tag struct {
	Id          int       `orm:"auto"`
	Name        string    `orm:"unique;type(text)"`
	CreateDate  time.Time `orm:"auto_now_add;type(datetime)"`
	Description string    `orm:"null;type(text)"`
	Link        []*Link   `orm:"null;reverse(many);on_delete(do_nothing)"`
}

func TagReadOrCreate(name string) (tag *Tag) {
	if name != "" {
		o := orm.NewOrm()
		tag = new(Tag)
		tag.Name = name
		if _, id, err := o.ReadOrCreate(tag, "Name"); err == nil {
			tag.Id = int(id)
		} else {
			beego.Error(err)
		}
	}
	return tag
}
func LinkReadOrCreate(link *Link, tags string) {
	o := orm.NewOrm()

	//save tags to db
	if tags != "" {
		tagArr := strings.Split(tags, ",")
		link.Tags = make([]*Tag, len(tagArr))
		for index, name := range tagArr {
			tag := TagReadOrCreate(name)
			if tag != nil {
				link.Tags[index] = tag
			}
		}
	}

	// 三个返回参数依次为：是否新创建的，对象Id值，错误
	if _, _, err := o.ReadOrCreate(link, "Url"); err == nil {
		if len(link.Tags) > 0 {
			m2m := o.QueryM2M(link, "Tags")
			if _, err := m2m.Add(link.Tags); err != nil {
				beego.Error(err)
			}
		}
	} else {
		beego.Error(err)
	}
}


//https://github.com/lyonlai/bootstrap-paginator
func LinkPage(p int, size int) utils.Page{
	o := orm.NewOrm()
	var link Link
	var list []Link
	qs := o.QueryTable(link)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-ModifyDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
func LinkDelete(link *Link) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.Delete(link); err == nil {
		if num>0{
			result=true
		}
	}
	return result
}

func ImportRemoteLinkTable(){
	o1 := orm.NewOrm()
	o1.Using("default")

	o2 := orm.NewOrm()
	o2.Using("remote")

	var localLinks []Link
	qs1 := o1.QueryTable("link")
	qs1.RelatedSel().All(&localLinks)

	var remoteLinks []Link
	qs2 := o2.QueryTable("link")
	qs2.RelatedSel().All(&remoteLinks)

	result := make([]Link,0)

	for _,v1:=range remoteLinks{
		have:=false
		for _,v2:=range localLinks{
			if v1.Url==v2.Url{
				have=true
				break
			}
		}

		if !have{
			v1.Id=0
			for _,t:=range v1.Tags{
				t.Id=0
				_, id, err := o1.ReadOrCreate(&t, "Name")
				if err == nil {
					t.Id=int(id)
				}else{
					beego.Error(err)
				}
			}
			result = append(result,v1)
		}
	}

	if len(result)>0{
		if _,err:=o1.InsertMulti(len(result),&result);err!=nil{
			beego.Error(err)
		}
	}
}

func init() {
	// 需要在init中注册定义的model
	//maxIdle := 30
	//maxConn := 30
	orm.RegisterModel(new(Link), new(Tag))
}
