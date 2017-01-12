package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"web/utils"
)

type Book struct {
	Id int			`orm:"auto"`
	Name string		`orm:"unique;type(text)"`
	Content string		`orm:"null;type(text)"`
	ChapterRules string	`orm:"null;type(text)"`
	ContentRules string	`orm:"null;type(text)"`
	Maker string		`orm:"null"`
	Chapters []*Chapter	`orm:"null;reverse(many);on_delete(cascade)"`
	Url string		`orm:"null;type(text)"`
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

type Chapter struct {
	Id int			`orm:"auto"`
	Index int		`orm:"null"`
	Title string		`orm:"null;type(text)"`
	Content string		`orm:"null;type(text)"`
	Book *Book		`orm:"rel(fk)"`
	CreateDate  time.Time 	`orm:"auto_now_add;type(datetime)"`
	ModifyDate  time.Time 	`orm:"auto_now;type(datetime)"`
}

func BookPage(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj Book
	var list []Book
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}
func ChapterPage(p int, size int) utils.Page{
	o := orm.NewOrm()
	var obj Book
	var list []Chapter
	qs := o.QueryTable(obj)
	count, _ := qs.Limit(-1).Count()
	qs.RelatedSel().OrderBy("-CreateDate").Limit(size).Offset((p - 1) * size).All(&list)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

func SaveBook(book *Book) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(book)
	return id
}

func SaveChapter(c *Chapter) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(c)
	return id
}

func BookSaveOrUpdate(book *Book) int64{
	o := orm.NewOrm()
	count:=int64(0)

	obj:=Book{Id:book.Id}
	if o.Read(&obj) == nil {
		if num, err := o.Update(&book); err == nil {
			count=num
		}
	}
	return count
}

func ChapterUpdate(chapter *Chapter) int64{
	o := orm.NewOrm()
	count:=int64(0)
	if o.Read(&chapter) == nil {
		if num, err := o.Update(&chapter); err == nil {
			count=num
		}
	}
	return count
}

func init() {
	orm.RegisterModel(new(Book), new(Chapter))
}
