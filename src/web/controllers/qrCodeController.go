package controllers

import "github.com/astaxie/beego"

type QrCodeController struct {
	beego.Controller
}
func (this *QrCodeController) URLMapping() {
	this.Mapping("/qr/index", this.Index)
	this.Mapping("/qr/decode", this.Decode)
}
//@router /qr/index [get]
func (this *QrCodeController) Index() {
	this.TplName="qrcode/index.tpl"
}
//@router /qr/decode [get]
func (this *QrCodeController) Decode() {
	this.TplName="qrcode/decode.tpl"
}