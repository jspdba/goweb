package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
	"strconv"
	"web/service"
	"web/cache"
)

type BookController struct {
	beego.Controller
}
//bm := NewBeeMap()
//go get github.com/astaxie/beego/cache

func (this *BookController) URLMapping() {
	this.Mapping("/book/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/book/save", this.SaveOrUpdate)
	this.Mapping("/book/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/book/list", this.List)
	this.Mapping("/book/taskUpdate/:id([0-9]{0,})", this.TaskUpdate)
	this.Mapping("/book/url/info", this.UrlInfo)
	this.Mapping("/book/search", this.Search)
}

// @router /book/edit/:id([0-9]{0,}) [get]
func (this *BookController) Edit() {
	id:=this.Ctx.Input.Param(":id")
	bk:= models.Book{}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,book:=models.FindBookById(i)
			if ok{
				bk=book
			}
		}
	}
	this.Data["entry"] = bk
	this.TplName = "book/edit.tpl"
}

// @router /book/save [post]
func (this *BookController) SaveOrUpdate() {
	book := models.Book{}
	if err := this.ParseForm(&book); err != nil {
		beego.Error(err)
	}
	models.BookSaveOrUpdate(&book)
	this.Redirect("/book/list", 302)
}
// @router /book/taskUpdate/:id([0-9]{0,}) [get]
func (this *BookController) TaskUpdate() {
	id:=this.Ctx.Input.Param(":id")
	json:=JsonObj{Code: -1, Msg:"error"}
	if id!=""{
		if i,err:=strconv.ParseInt(id, 10, 64); err==nil{
			ok,book:=models.FindBookById(i)
			if ok{
				/*tag:="cache_book_"+id
				if !cache.IsExist(tag){
					cache.Put(tag,time.Now().Format("2006-01-02 15:04:05"),time.Second*60*10)
					//更新章节
					go (func(book *models.Book) bool{
						if book!=nil{
							//申请任务调度
							if book.Url!=""{
								chapters:= service.GetUrlInfo(book.Url,book.ChapterRules,-1)
								//增加index
								for i:=len(chapters);i>0;i--{
									chapters[i-1].Index=i
								}
								if ok,ch:=models.FindMaxIndexChapter(book);ok{
									if ch.Index<len(chapters){
										chapters=chapters[ch.Index:]
									}else{
										cache.Delete(tag)
										return false
									}
								}

								service.GetChapterContent(book.ContentRules,chapters,100)
								for _,chapter := range chapters{
									chapter.Book= book
								}
								beego.Info("begin >>>>>>>>>>>>")
								models.ChapterInsertMulti(chapters)
								beego.Info("<<<<<<<<<<< over")
							}
						}
						cache.Delete(tag)
						return true
					})(&book)
					json = JsonObj{Code: 0, Msg:"ok"}
				}else{
					json = JsonObj{Code: -1, Msg:tag+"=正在更新！更新时间="+cache.Get(tag).(string)}
				}*/
				tag:="cache_book_"+id
				if isRunning:=service.UpdateBook(&book,tag);isRunning{
					json = JsonObj{Code: -1, Msg:tag+"=正在更新！更新时间="+cache.Get(tag).(string)}
				}else{
					json = JsonObj{Code: 0, Msg:"ok"}
				}
			}
		}
	}
	this.Data["json"] = json
	this.ServeJSON()
}

// @router /book/delete/:id([0-9]+)
func (this *BookController) Delete(){
	id:=this.Ctx.Input.Param(":id")
	if id!=""{
		if i,err:=strconv.Atoi(id); err==nil{
			book := models.Book{Id:i}
			models.BookDelete(&book)
		}

	}
}
// @router /book/url/info
func (this *BookController) UrlInfo(){
	url:=this.GetString("Url")
	book:=&models.Book{}
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"no result",
		"result":nil,
	}
	if url!=""{
		book=service.GetBookInfo(url)
		jsonMap=map[string]interface{}{
			"code":0,
			"msg":"",
			"result":book,
		}
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}

// @router /book/search
func (this *BookController) Search(){
	name:=this.GetString("Name")
	book:=&models.Book{}
	jsonMap:=map[string]interface{}{
		"code":-1,
		"msg":"no result",
		"result":nil,
	}
	if name!=""{
		book=service.Search(name)
		jsonMap=map[string]interface{}{
			"code":0,
			"msg":"",
			"result":book,
		}
	}

	this.Data["json"] = jsonMap
	this.ServeJSON()
}

// @router /book/list
func (this *BookController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.BookPage(page.PageNo,page.PageSize)
	this.TplName = "book/list.tpl"
}