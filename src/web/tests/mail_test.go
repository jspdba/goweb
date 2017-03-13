package test

import (
	"testing"
	"time"
	"web/utils"
	"fmt"
)

func Test_mail(t *testing.T) {
	err := utils.SendMail("jspdba@163.com", "***", "smtp.163.com:25", "jspdba@163.com", "for test", "点击这里修改密码：http://www.xxx.com/updatepass?&time="+time.Now().Format("2006-01-0215:04:05"), "html")
	if err!=nil{
		fmt.Println(err.Error())
	}
}