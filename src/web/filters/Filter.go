package filters

import (
    "github.com/astaxie/beego/context"
    "github.com/astaxie/beego"
    "web/models"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
    token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
    var user models.User
    if flag {
        flag, user = models.FindUserByToken(token)
    }
    return flag, user
}

var FilterUser = func(ctx *context.Context) {
    ok, _ := IsLogin(ctx)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}
