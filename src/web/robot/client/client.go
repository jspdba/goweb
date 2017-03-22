package main

import (
	"errors"
	"net/http"
	"github.com/juju/persistent-cookiejar"
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"io/ioutil"
	"strings"
	"fmt"
	"os"
	"path"
	"encoding/json"
)

const (
	userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36"
	connection     = "keep-alive"
	//pragma         = "no-cache"
	xhr            = "XMLHttpRequest"
	acceptEncoding = ""
	accept = "application/json, text/javascript, */*; q=0.01"
	acceptLanguage ="zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4,nb;q=0.2"
	origin ="http://bjguahao.gov.cn"
)

var (
	host   = "bjguahao.gov.cn"
	AuthError = errors.New("验证返回错误")
)

type User struct {
	host   string
	Client *http.Client
	Config *conf
	headers *http.Header
}
//配置文件读取
type conf struct {
	Username string		`json:"username"`
	Password string		`json:"password"`
}
//xhr结果
type Result struct {
	Data interface{}	`json:"data"`
	HasError bool		`json:"hasError"`
	Code int		`json:"code"`
	Msg string		`json:"msg"`
}

func (this *User) Init() {
	cookieJar, _ := cookiejar.New(nil)

	if this.Client == nil {
		this.Client = &http.Client{
			Jar:       cookieJar,
		}
	}
	this.LoadConfig("")

	headers:=newHTTPHeaders(true)
	this.headers = &headers
}

func(this *User) LoadConfigWithBeego(){
	if this.Config==nil{
		this.Config = new(conf)
		this.Config.Username = beego.AppConfig.String("username")
		this.Config.Password = beego.AppConfig.String("password")
	}
}

func (this *User) LoadConfig(cfg string) {
	if cfg==""{
		dir,e:=os.Getwd()
		if e!=nil{
			return
		}
		cfg=path.Join(dir,"src/web/guahao.json")
	}

	fd, err := os.Open(cfg)
	if err != nil {
		panic("无法打开配置文件 config.json: " + err.Error())
	}
	defer fd.Close()

	config := new(conf)
	err = json.NewDecoder(fd).Decode(&config)
	if err != nil {
		panic("解析配置文件出错: " + err.Error())
	}
	//读取配置文件，且放到user中
	this.Config = config
}

func (this *User) SaveCookie() {
	this.Client.Jar.(*cookiejar.Jar).Save()
}

func newHTTPHeaders(isXhr bool) http.Header {
	headers := make(http.Header)
	headers.Set("Accept", accept)
	headers.Set("User-Agent", userAgent)
	headers.Set("Host", host)
	headers.Set("Connection", connection)
	//headers.Set("Accept-Encoding", acceptEncoding)
	headers.Set("Accept-Language", acceptLanguage)
	headers.Set("Origin", origin)
	headers.Set("Referer", "http://bjguahao.gov.cn/logout.htm")
	//headers.Set("Pragma", pragma)
	if isXhr {
		headers.Set("X-Requested-With", xhr)
	}
	return headers
}


func (this *User) Post(path, contentType string, content io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", path, content)
	if err != nil {
		return
	}

	this.headers.Set("Content-Type", contentType)
	req.Header = *this.headers
	resp, err = this.Client.Do(req)
	if err != nil {
		beego.Error(err)
	}
	return
}

func(this *User) Login(){
	var path = "http://bjguahao.gov.cn/quicklogin.htm"

	v := url.Values{}
	v.Set("mobileNo", this.Config.Username)
	v.Set("password", this.Config.Password)
	v.Set("isAjax", "true")
	v.Set("yzm", "")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	contentType := "application/x-www-form-urlencoded; charset=UTF-8"
	resp, err:=this.Post(path,contentType,body)
	if err!=nil{
		beego.Error(err)
		return
	}
	this.SaveCookie()
	err1,b:=CheckResp(resp,err)
	if err1!=nil{
		beego.Error(err1)
		return
	}
	decodeResult(b)
}

func(this *User) QueryDoctor(){
	var path = "http://bjguahao.gov.cn/dpt/partduty.htm"

	v := url.Values{}
	v.Set("hospitalId", "114")
	v.Set("departmentId", "200000802")
	v.Set("dutyCode", "1")
	v.Set("dutyDate", "true")
	v.Set("isAjax", "2017-03-24")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	contentType := "application/x-www-form-urlencoded; charset=UTF-8"
	this.headers.Set("Referer","http://bjguahao.gov.cn/dpt/appoint/114-200000802.htm")
	resp, err:=this.Post(path,contentType,body)
	if err!=nil{
		beego.Error(err)
		return
	}
	err1,b:=CheckResp(resp,err)
	if err1!=nil{
		beego.Error(err1)
		return
	}
	decodeResult(b)
}

func decodeResult(res []byte) (bool,string){
	config := new(Result)
	err := json.NewDecoder(strings.NewReader(string(res))).Decode(&config)
	if err != nil {
		panic("解析配置文件出错: " + err.Error())
		return false,err.Error()
	}

	if config.HasError{
		beego.Error(config.Msg)
		return false,config.Msg
	}else{
		beego.Info(config.Msg)
		return true,""
	}
}

func CheckResp(resp *http.Response, err error) (err1 error, s []byte) {
	if err != nil {
		beego.Error("Get Error：" + err.Error())
		return err, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("State Error，StatusCode = %d", resp.StatusCode))
		return AuthError, nil
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("读取响应内容失败：%s", err.Error())
		return err, nil
	}
	beego.Info("相应内容=%s", string(content))
	return nil, content
}

func main() {
	var user=new(User)
	user.Init()
	user.Login()
	user.QueryDoctor()
}