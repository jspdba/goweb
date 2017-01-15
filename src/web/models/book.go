package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
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
	Url string		`orm:"null;type(text)"`
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

func ChapterPage(p int, size int,bookId int) utils.Page{
	o := orm.NewOrm()
	var obj Chapter
	var list []Chapter
	qs := o.QueryTable(obj)
	count:=int64(0)
	if bookId>=0{
		count, _ = qs.Filter("Book__Id", bookId).Limit(-1).Count()
		qs.RelatedSel().OrderBy("index").Filter("Book__Id", bookId).Limit(size).Offset((p - 1) * size).All(&list)
	}else{
		count, _ = qs.Limit(-1).Count()
		qs.RelatedSel().OrderBy("index").Limit(size).Offset((p - 1) * size).All(&list)
	}

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


func ChapterInsertMulti(chapters []*Chapter){
	o := orm.NewOrm()
	if _, err :=o.InsertMulti(len(chapters),chapters);err!=nil{
		beego.Error(err)
	}
}

func ChapterSaveOrUpdate(chapter *Chapter) int64{
	o := orm.NewOrm()

	chapterOld:=*chapter
	if created, id, err := o.ReadOrCreate(chapter, "Id"); err == nil {
		if created {
			return id
		} else {
			chapter.Title = chapterOld.Title
			chapter.Url = chapterOld.Url
			chapter.Content = chapterOld.Content
			chapter.Book = chapterOld.Book
			chapter.Content = chapterOld.Content
			chapter.Index = chapterOld.Index

			if id, err := o.Update(chapter); err==nil{
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
func ChapterNext(chapter *Chapter) Chapter{
	o := orm.NewOrm()
	var obj Chapter
	qs := o.QueryTable(obj)
	qs.Filter("book", chapter.Book).Filter("index__gt",chapter.Index).OrderBy("index").Limit(1).One(&obj)
	return obj
}
func ChapterPre(chapter *Chapter) Chapter{
	o := orm.NewOrm()
	var obj Chapter
	qs := o.QueryTable(obj)
	qs.Filter("book", chapter.Book).Filter("index__lt",chapter.Index).OrderBy("-index").Limit(1).One(&obj)
	return obj
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
func ChapterDelete(chapter *Chapter) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.Delete(&chapter); err == nil {
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
func FindChapterById(id int64) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("Id", id).One(&entity)
	return err != orm.ErrNoRows, entity
}
//根据主键查找chapter
func FindChapter(chapter *Chapter) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("Id", chapter.Id).RelatedSel().One(&entity)
	return err != orm.ErrNoRows, entity

}

func init() {
	orm.RegisterModel(new(Book), new(Chapter))
}
