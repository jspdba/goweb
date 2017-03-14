//登录且下载pdf文件
package downLoad

import (
	"net/http"
	"github.com/juju/persistent-cookiejar"
	"log"
)

type User struct {
	LoginId string `json:"loginId"`
	Password string `json:"password"`
	Client *http.Client
}


func (this *User) Init() {
	cookieJar, _ := cookiejar.New(nil)

	if this.Client == nil {
		this.Client = &http.Client{
			Jar:       cookieJar,
		}
	}
}

func(this *User) Login()  {

}

func(this *User) LinkUrl() {
	var url string="https://m.aliyun.com/yunqi/articles/69316?utm_campaign=zilxiaz&utm_medium=images&utm_source=renyimen&utm_content=m_12058"
	log.Println(url)
}