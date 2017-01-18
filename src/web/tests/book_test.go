package test

import (
	"testing"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"web/models"
	"strings"
	"log"
)

//测试读取数据库
//读取章节：）
func Test_book_chapter(t *testing.T) {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/beego?charset=utf8", 30)
	//orm.RunSyncdb("default", false, true)

	/*page:=models.BookPage(1,10)
	list:=page.List

	for _,book:= range list.([]models.Book){
		beego.Info(book)
	}*/
	if ok,chapter:=models.FindChapterByBookIdAndIndex(1,369); ok{
		content:=addLine(chapter.Content)
		for _,str:= range strings.Split(content,"\r\n"){
			log.Println(str)
		}
	}
}


func addLine(str string) string{
	str=strings.Replace(str,"&nbsp;&nbsp;&nbsp;&nbsp;","\r\n&nbsp;&nbsp;&nbsp;&nbsp;",-1)
	str=strings.Replace(str,"    ","\r\n    ",-1)
	return str
}
