package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) URLMapping() {
	this.Mapping("/user/list)", this.List)
}

//验证登录
func (c *LoginController) Login() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if flag, user := models.Login(username, password); flag {
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), user.Token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	} else {
		flash.Error("用户名或密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/login", 302)
	}
}

 //@router /user/list
func (this *LoginController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["page"] = models.Page(page.PageNo,page.PageSize)
	this.TplName = "user/list.tpl"
}

