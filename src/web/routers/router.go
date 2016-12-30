package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.LinkController{})
	beego.Include(&controllers.LoginController{})
}
