package controllers

import (
	"github.com/astaxie/beego"
	"web/models"
	"web/utils"
	"github.com/sluu99/uuid"
	"crypto/md5"
	"encoding/hex"
	"web/filters"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) URLMapping() {
	this.Mapping("/user/list)", this.List)
	this.Mapping("/logout)", this.Logout)
}
//登录页
//@router /login [get]
func (c *UserController) LoginPage() {
	IsLogin, _ := filters.IsLogin(c.Ctx)
	if IsLogin {
		c.Redirect("/", 302)
	} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["title"] = "登录"
		c.TplName = "user/login.tpl"
	}
}

//验证登录
//@router /login [post]
func (c *UserController) Login() {
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
//注册页
// @router /register [get]
func (c *UserController) RegisterPage() {
	//IsLogin, _ := filters.IsLogin(c.Ctx)
	//if IsLogin {
	//	c.Redirect("/", 302)
	//} else {
		beego.ReadFromRequest(&c.Controller)
		c.Data["title"] = "注册"
		c.TplName = "user/register.tpl"
	//}
}
//验证注册
// @router /register [post]
func (c *UserController) Register() {
	flash := beego.NewFlash()
	username, password := c.Input().Get("username"), c.Input().Get("password")
	if len(username) == 0 || len(password) == 0 {
		flash.Error("用户名或密码不能为空")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else if flag, _ := models.FindUserByUserName(username); flag {
		flash.Error("用户名已被注册")
		flash.Store(&c.Controller)
		c.Redirect("/register", 302)
	} else {
		var token = uuid.Rand().Hex()
		h := md5.New()
		h.Write([]byte(password))
		user := models.User{Username: username, Password: hex.EncodeToString(h.Sum(nil)), Avatar: "/static/imgs/avatar.png", Token: token}
		models.SaveUser(&user)
		// others are ordered as cookie's max age time, path,domain, secure and httponly.
		c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), token, 30 * 24 * 60 * 60, "/", beego.AppConfig.String("cookie.domain"), false, true)
		c.Redirect("/", 302)
	}
}

//登出
//@router /logout
func (c *UserController) Logout() {
	c.SetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"), "", -1, "/", beego.AppConfig.String("cookie.domain"), false, true)
	c.Redirect("/", 302)
}

 //@router /user/list
func (this *UserController) List() {
	page:=utils.Page{PageNo:1,PageSize:20}
	if err := this.ParseForm(&page); err != nil {
		beego.Error(err)
	}
	this.Data["title"] = "用户列表"
	this.Data["page"] = models.Page(page.PageNo,page.PageSize)
	this.Layout=""
	this.TplName = "user/list.tpl"
}

