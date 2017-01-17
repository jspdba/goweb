package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
)

type ChapterLogController struct {
	beego.Controller
}

func (this *ChapterLogController) URLMapping() {
	this.Mapping("/log/chapter/:tag([\\w]+)/:bookId([0-9]+)/:index([0-9]+)", this.Add)
	this.Mapping("/log/tag/:tag([\\w]+)", this.FindTag)
}
// @router /log/chapter/:tag([\w]+)/:bookId([0-9]+)/:index([0-9]+) [get]
func (this *ChapterLogController) Add(){
	log := models.Log{}
	resMap:=map[string]interface{}{
		"code":0,
		"msg":"",
		"data":nil,
	}
	if this.Ctx.Input.Param(":tag")==""{
		resMap=map[string]interface{}{
			"code":-1,
			"msg":"tag 参数错误",
			"data":nil,
		}

	}else  if this.Ctx.Input.Param(":bookId")==""{

		resMap=map[string]interface{}{
			"code":-1,
			"msg":"bookId 参数错误",
			"data":nil,
		}


	}else if this.Ctx.Input.Param(":index")==""{
		resMap=map[string]interface{}{
			"code":-1,
			"msg":"index 参数错误",
			"data":nil,
		}

	}else{
		log.Tag = this.Ctx.Input.Param(":tag")
		log.Index = this.Ctx.Input.Param(":index")
		log.BookId = this.Ctx.Input.Param(":bookId")
		if ok:=models.LogInsertOrUpdate(&log) ;ok{
			resMap=map[string]interface{}{
				"code":0,
				"msg":"ok",
				"data":nil,
			}
		}
	}

	this.Data["json"]=resMap
	this.ServeJSON()
}

 // @router /log/tag/:tag([\w]+) [get]
func (this *ChapterLogController) FindTag(){
	resMap:=map[string]interface{}{
		"code":0,
		"msg":"",
		"result":0,
	}
	tag := this.Ctx.Input.Param(":tag")
	count:=models.FindLogCount(tag)
	resMap["result"]=count;
	this.Data["json"]=resMap
	this.ServeJSON()
}