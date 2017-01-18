package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["web/controllers:BookController"] = append(beego.GlobalControllerRouter["web/controllers:BookController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/book/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:BookController"] = append(beego.GlobalControllerRouter["web/controllers:BookController"],
		beego.ControllerComments{
			Method: "SaveOrUpdate",
			Router: `/book/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:BookController"] = append(beego.GlobalControllerRouter["web/controllers:BookController"],
		beego.ControllerComments{
			Method: "TaskUpdate",
			Router: `/book/taskUpdate/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:BookController"] = append(beego.GlobalControllerRouter["web/controllers:BookController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/book/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:BookController"] = append(beego.GlobalControllerRouter["web/controllers:BookController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/book/list`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Edit",
			Router: `/chapter/edit/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Detail",
			Router: `/chapter/detail/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Next",
			Router: `/chapter/next/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Pre",
			Router: `/chapter/pre/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "SaveOrUpdate",
			Router: `/chapter/save`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/chapter/delete/:id([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/chapter/list/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "HasNewChapter",
			Router: `/chapter/new/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "ListByLog",
			Router: `/chapter/list/:tag(\w+)/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterController"],
		beego.ControllerComments{
			Method: "FindByTitle",
			Router: `/chapter/title/:id([0-9]{0,})`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterLogController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterLogController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/log/chapter/:tag([\w]+)/:bookId([0-9]+)/:index([0-9]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["web/controllers:ChapterLogController"] = append(beego.GlobalControllerRouter["web/controllers:ChapterLogController"],
		beego.ControllerComments{
			Method: "FindTag",
			Router: `/log/tag/:tag([\w]+)`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

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

	beego.GlobalControllerRouter["web/controllers:QrCodeController"] = append(beego.GlobalControllerRouter["web/controllers:QrCodeController"],
		beego.ControllerComments{
			Method: "Index",
			Router: `/qr/index`,
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
