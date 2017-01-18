package controllers

import "github.com/astaxie/beego"

type QrCodeController struct {
	beego.Controller
}
func (this *QrCodeController) URLMapping() {
	this.Mapping("/qr/index", this.Index)
}
//@router /qr/index [get]
func (this *QrCodeController) Index() {
	this.TplName="qrcode/index.tpl"
}