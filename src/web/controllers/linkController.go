package controllers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"strings"
	"web/models"
)

type LinkController struct {
	beego.Controller
}

type Data struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type JsonObj struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data *Data  `json:"data"`
}

func (this *LinkController) URLMapping() {
	this.Mapping("/link/edit/:id([0-9]+)", this.Edit)
	this.Mapping("/link/save", this.Save)
	this.Mapping("/link/delete/:id([0-9]+)", this.Delete)
	this.Mapping("/link/list", this.List)
	this.Mapping("/link/info", this.Info)
}

// @router /link/edit/:id([0-9]{0,}) [get]
func (this *LinkController) Edit() {
	this.TplName = "link/edit.tpl"
}

// @router /link/save [post]
func (this *LinkController) Save() {
	link := models.Link{}
	if err := this.ParseForm(&link); err != nil {
		beego.Error(err)
	}
	tags := this.GetString("Tags.Name")
	models.LinkReadOrCreate(&link, tags)
	this.Redirect("/link/list", 302)
}

// @router /link/delete/:id
func (this *LinkController) Delete() {
}

// @router /link/list
func (this *LinkController) List() {
	this.TplName = "link/list.tpl"
}

// @router /link/info
func (this *LinkController) Info() {
	url := strings.Trim(this.GetString("url"), "")
	if url == "" {
		this.Data["json"] = &JsonObj{Code: 1, Msg: "参数错误，url=nil"}
	} else {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		beego.Info(url)
		this.Data["json"] = JsonObj{Code: 0, Data: getUrlInfo(url)}
	}
	this.ServeJSON()
}

func getUrlInfo(url string) (data *Data) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		beego.Error(err)
		return &Data{}
	}
	if doc != nil {
		title := doc.Find("title").Text()
		content, _ := doc.Find("meta[name=description]").Attr("content")
		data = &Data{title, content}
	}
	return data
}

/*func getUrlInfo(url string) (data Data){
	req:=httplib.Get(url).Debug(true).SetTimeout(10 * time.Second, 3 * time.Second)
	if url!=""{
		if strings.HasPrefix(url,"https://"){
			req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		}
	}
	result,error:=req.String()
	if error!=nil{
		beego.Error(error)
	}
	beego.Info(result)
	return data
}*/
