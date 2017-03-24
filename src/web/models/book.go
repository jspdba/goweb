package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"web/utils"
	"errors"
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
func ChapterPageByLog(p int, size int,log *Log) utils.Page{
	o := orm.NewOrm()
	var obj Chapter
	var list []Chapter
	qs := o.QueryTable(obj)
	count:=int64(0)

	bookId,err:=strconv.Atoi(log.BookId)
	if err==nil{
		if bookId>=0{
			count, _ = qs.Filter("index__gte",log.Index).Filter("Book__Id", bookId).Limit(-1).Count()
			qs.RelatedSel().OrderBy("index").Filter("index__gte",log.Index).Filter("Book__Id", bookId).Limit(size).Offset((p - 1) * size).All(&list)
		}else{
			count, _ = qs.Limit(-1).Count()
			qs.RelatedSel().OrderBy("index").Limit(size).Offset((p - 1) * size).All(&list)
		}
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
	//分批插入，每次max
	max:=1000
	length:=len(chapters)
	if length>0{
		if length<=max{
			if _, err :=o.InsertMulti(length,chapters);err!=nil{
				beego.Error(err)
			}
		}else{
			itemcount:=length/max
			for i,start,end:=0,0,max;i<itemcount;i++{
				cha:=chapters[start:end]
				if _, err :=o.InsertMulti(len(cha),cha);err!=nil{
					beego.Error(err)
				}
				start=end
				end+=max
			}
			if(length%max!=0){
				cha:=chapters[length-length%max:]
				if _, err :=o.InsertMulti(len(cha),cha);err!=nil{
					beego.Error(err)
				}
			}
		}
	}

}

func ChapterSaveOrUpdate(chapter *Chapter) int64{
	o := orm.NewOrm()

	chapterOld:=*chapter
	if created, id, err := o.ReadOrCreate(chapter, "Id",); err == nil {
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
		if num, err := o.Update(chapter); err == nil {
			return num
		}
	return int64(0)
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

func ChapterDeleteBook(id int) bool{
	o := orm.NewOrm()
	result:=false
	if num, err := o.QueryTable("chapter").Filter("book__id__eq", id).Delete(); err == nil {
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

func FindBookByStrId(id string) (bool, Book) {
	o := orm.NewOrm()
	var book Book
	if i, err := strconv.Atoi(id); err == nil {
		err := o.QueryTable(book).Filter("Id", i).One(&book)
		return err != orm.ErrNoRows, book
	}
	return false,book
}
func FindChapterById(id int64) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("Id", id).One(&entity)
	return err != orm.ErrNoRows, entity
}
func FindChapterByStrId(id string) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	if i,er:=strconv.Atoi(id);er==nil{
		err := o.QueryTable(entity).Filter("Id", i).RelatedSel().One(&entity)
		return err != orm.ErrNoRows, entity
	}
	return false,entity
}
//根据主键查找chapter
func FindChapter(chapter *Chapter) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("Id", chapter.Id).RelatedSel().One(&entity)
	return err != orm.ErrNoRows, entity

}

func FindMaxIndexChapter(book *Book) (bool, *Chapter){
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("book__id",book.Id).OrderBy("-index").Limit(1).One(&entity)
	return err != orm.ErrNoRows, &entity
}
func FindMaxIndexChapterByBookId(bookId string) (bool, *Chapter){
	o := orm.NewOrm()
	var entity Chapter
	if bookId==""{
		return false,nil
	}

	id,e:=strconv.Atoi(bookId);
	if e!=nil{
		return false,nil
	}
	err := o.QueryTable(entity).RelatedSel().Filter("book__id",id).OrderBy("-index").Limit(1).One(&entity)
	return err != orm.ErrNoRows, &entity
}

func ChapterPageByTitle(p int, size int,title string,bid string) utils.Page{
	o := orm.NewOrm()
	var obj Chapter
	var list []Chapter
	qs := o.QueryTable(obj)
	count:=int64(0)

	bookId,err:=strconv.Atoi(bid)
	if err==nil{
		if bookId>=0{
			count, _ = qs.Filter("title__contains",title).Filter("Book__Id", bookId).Limit(-1).Count()
			qs.RelatedSel().OrderBy("index").Filter("title__contains",title).Filter("Book__Id", bookId).Limit(size).Offset((p - 1) * size).All(&list)
		}else{
			count, _ = qs.Filter("title__contains",title).Limit(-1).Count()
			qs.RelatedSel().OrderBy("index").Filter("title__contains",title).Limit(size).Offset((p - 1) * size).All(&list)
		}
	}

	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.PageUtil(c, p, size, list)
}

//根据主键查找chapter
func FindChapterByBookIdAndIndex( bookId int, index int) (bool, Chapter) {
	o := orm.NewOrm()
	var entity Chapter
	err := o.QueryTable(entity).Filter("book__id",bookId).Filter("Index", index).RelatedSel().One(&entity)
	return err != orm.ErrNoRows, entity
}
func ImportRemoteBookTable(){
	o1 := orm.NewOrm()
	o1.Using("default")

	o2 := orm.NewOrm()
	o2.Using("remote")

	var localBooks []Book
	qs1 := o1.QueryTable("book")
	qs1.RelatedSel().All(&localBooks)

	var remoteBooks []Book
	qs2 := o2.QueryTable("book")
	qs2.RelatedSel().All(&remoteBooks)

	result := make([]Book,0)

	for _,v1:=range remoteBooks{
		have:=false
		for _,v2:=range localBooks{
			if v1.Url==v2.Url{
				have=true
				break
			}
		}

		if !have{
			v1.Id=0
			result = append(result,v1)
		}
	}

	if len(result)>0{
		if _,err:=o1.InsertMulti(len(result),&result);err!=nil{
			beego.Error(err)
		}
	}
}

func Export(id string) error{
	//this.obj
	localOrm := orm.NewOrm()
	localOrm.Using("default")

	remoteOrm := orm.NewOrm()
	remoteOrm.Using("remote")

	//查询远程是否有这本书

	var localBook Book
	err:=localOrm.QueryTable("book").Filter("Id",id).One(&localBook)
	if err!=nil{
		beego.Error(err)
		return errors.New("参数错误！book.id="+id)
	}

	name:=localBook.Name

	var remoteBook Book
	err=remoteOrm.QueryTable("book").Filter("Name",name).One(&remoteBook)
	if err!=nil{
		beego.Error(err)
		return errors.New("服务器不存在这本书="+name)
	}

	//1.更新远程空章节，2.插入远程缺失的章节
	var localChapterList,remoteChapterList []*Chapter
	_,err = remoteOrm.QueryTable("chapter").Filter("Book__Id__eq", remoteBook.Id).Filter("Content","").OrderBy("Index").All(&remoteChapterList)
	if err!=nil{
		beego.Error(err)
		return errors.New("远程查询空章节失败="+remoteBook.Name)
	}

	if len(remoteChapterList)>0{
		var indexs = []int{}
		for _,chapter := range remoteChapterList{
			indexs = append(indexs,chapter.Index)
		}
		//查询本地数据
		_,err = remoteOrm.QueryTable("chapter").Filter("Book__Id__eq", localBook.Id).Filter("Index__in",indexs).OrderBy("Index").All(&localChapterList)

		//更新content值
		for _,localChapter:=range localChapterList{
			for _,remoteChapter:=range remoteChapterList{
				if localChapter.Index==remoteChapter.Index {
					remoteChapter.Content=localChapter.Content
					break
				}else if remoteChapter.Index>localChapter.Index{
					break
				}
			}
		}
		//更新远程数据库
		p, err := remoteOrm.Raw("UPDATE chapter SET content = ? WHERE id = ?").Prepare()
		if err!=nil{
			beego.Error(err)
			p.Close()
			return err
		}
		for _,chapter:=range remoteChapterList{
			if chapter.Content!=""{
				_, err := p.Exec(chapter.Content, chapter.Id)
				if err!=nil{
					beego.Error(err)
				}
			}
		}
		p.Close()
	}


	//增加远程数据
	var remoteChapterOfMaxIndex Chapter
	err = remoteOrm.Raw("SELECT max(index) index FROM user WHERE book_id = ?", remoteBook.Id).QueryRow(&remoteChapterOfMaxIndex)

	if err!=nil{
		beego.Error(err)
		return err
	}

	var toExportChapterList []*Chapter
	_,err=localOrm.QueryTable("chapter").Filter("Book__Id__eq",localBook.Id).Filter("Index__gt",remoteChapterOfMaxIndex.Index).All(&toExportChapterList)

	if err!=nil{
		beego.Error(err)
		return err
	}

	qs := remoteOrm.QueryTable("chapter")
	i, _ := qs.PrepareInsert()

	for _, chapter := range toExportChapterList {
		chapter.Id = 0

		chapter.Book.Id=remoteBook.Id

		_, err := i.Insert(chapter)
		if err != nil {
			beego.Error(err)
		}
	}
	i.Close()
	return err
}


func init() {
	orm.RegisterModel(new(Book), new(Chapter))
}
