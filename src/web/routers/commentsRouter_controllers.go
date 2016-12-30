package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["web/controllers:LinkController"] = append(beego.GlobalControllerRouter["web/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/link/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:LinkController"] = append(beego.GlobalControllerRouter["web/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Save",
			Router: `/link/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:LinkController"] = append(beego.GlobalControllerRouter["web/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/link/delete/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:LinkController"] = append(beego.GlobalControllerRouter["web/controllers:LinkController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/link/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:LinkController"] = append(beego.GlobalControllerRouter["web/controllers:LinkController"],
		beego.ControllerComments{
			Method: "Info",
			Router: `/link/info`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "LoginPage",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "RegisterPage",
			Router: `/register`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:UserController"] = append(beego.GlobalControllerRouter["web/controllers:UserController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/user/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
