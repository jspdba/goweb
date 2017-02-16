package service

import (
	"web/models"
	"github.com/astaxie/beego"
	"time"
	"web/cache"
)

func UpdateBook(book *models.Book,tag string) (running bool){
	if !cache.IsExist(tag){
		cache.Put(tag,time.Now().Format("2006-01-02 15:04:05"),time.Second*60*10)
		//更新章节
		go (func(book *models.Book) bool{
			if book!=nil{
				//申请任务调度
				if book.Url!=""{
					chapters:= GetUrlInfo(book.Url,book.ChapterRules,-1)
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

					GetChapterContent(book.ContentRules,chapters,100)
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
		})(book)
		running=false
	}else{
		running=true
	}
	return running
}
