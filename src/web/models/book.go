package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"web/utils"
	"github.com/astaxie/beego"
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

func BookSave(book *Book) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(book)
	return id
}

func ChapterSave(c *Chapter) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(c)
	return id
}

func ChapterSaveMulti(c [] *Chapter) int64 {
	o := orm.NewOrm()
	successNums, err:= o.InsertMulti(len(c),c)
	if err!=nil{
		beego.Error(err)
	}
	return successNums
}

func BookSaveOrUpdate(book *Book) int64{
	o := orm.NewOrm()
	/*err:=o.Read(&book,"Id")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")

		if id, err := o.Insert(book); err==nil{
			return id
		}else{
			beego.Error(err)
		}
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")

		if id, err := o.Insert(book); err==nil{
			return id
		}else{
			beego.Error(err)
		}
	} else {
		if id, err := o.Update(book); err==nil{
			return id
		}else{
			beego.Error(err)
		}
	}*/

	bookold:=*book
	if created, id, err := o.ReadOrCreate(book, "Id"); err == nil {
		if created {
			return id
		} else {
			book.Name = bookold.Name
			book.Url = bookold.Url
			book.ChapterRules = bookold.ChapterRules
			book.Chapters = bookold.Chapters
			book.Content = bookold.Content
			book.ContentRules = bookold.ContentRules
			book.Maker = bookold.Maker
			if id, err := o.Update(book); err==nil{
				return id
			}else{
				beego.Error(err)
			}
		}
	}
	return int64(-1)
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

func BookDelete(book *Book) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.Delete(&book); err == nil {
		if num>0{
			result=true
		}
	}
	return result
}
func FindBookById(id int64) (bool, Book) {
	o := orm.NewOrm()
	var book Book
	err := o.QueryTable(book).Filter("Id", id).One(&book)
	return err != orm.ErrNoRows, book
}
func init() {
	orm.RegisterModel(new(Book), new(Chapter))
}
