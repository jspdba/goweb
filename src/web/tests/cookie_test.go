package test

import (
	"net/http"
	"fmt"
	"net/http/cookiejar"
	"io/ioutil"
	"testing"
	"github.com/astaxie/beego"
	"time"
	"net"
	"net/url"
	"web/utils"
)
//get url response html
func GetUrlRespHtml(url string, myCookieJar *cookiejar.Jar,myCookieArray []*http.Cookie) (respHtml string ){
	fmt.Printf("getUrlRespHtml, url=%s", url)

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           myCookieJar,
	}
	httpReq, err := http.NewRequest("GET", url, nil)

	if err!=nil{
		beego.Error(err)
		return
	}

	httpReq.Header.Add("User-Agent", "android_gongjiao;v2.0.0;")

	httpResp, err := httpClient.Do(httpReq)
	if err!=nil{
		beego.Error(err)
		return
	}

	fmt.Printf("httpResp.Header=%s", httpResp.Header)
	fmt.Printf("httpResp.Status=%s", httpResp.Status)

	defer httpResp.Body.Close()

	if httpResp.StatusCode != 200{
		beego.Info(httpResp.StatusCode)
		return
	}
	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		fmt.Printf("get response for url=%s got error=%s\n", url, errReadAll.Error())
		beego.Error(errReadAll)
		return
	}

	if myCookieJar!=nil{
		respHtml = string(body)
		beego.Info(respHtml)
		myCookieArray = myCookieJar.Cookies(httpReq.URL)
	}
	return
}

func HttpPostForm(url1 string, data *url.Values,myCookieJar * cookiejar.Jar,myCookieArray []*http.Cookie) (content string){
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
		CheckRedirect: nil,
		Jar:           myCookieJar,
	}
	resp, err := client.PostForm(url1,*data)
	if err != nil {
		beego.Error(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
			return
		}
		content = string(body)
	}
	if myCookieJar!=nil{
		myCookieArray = myCookieJar.Cookies(resp.Request.URL)
		beego.Info(myCookieJar)
		beego.Info(len(myCookieArray))
		//PrintCurCookies(myCookieArray)
	}
	return

}
func PrintCurCookies(myCookieArray []*http.Cookie) {
	var cookieNum int = len(myCookieArray)
	fmt.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = myCookieArray[i]
		fmt.Printf("\n------ Cookie [%d]------", i)
		fmt.Printf("\tName=%s", curCk.Name)
		fmt.Printf("\tValue=%s", curCk.Value)
		fmt.Printf("\tPath=%s", curCk.Path)
		fmt.Printf("\tDomain=%s", curCk.Domain)
		fmt.Printf("\tExpires=%s", curCk.Expires)
		fmt.Printf("\tRawExpires=%s", curCk.RawExpires)
		fmt.Printf("\tMaxAge=%d", curCk.MaxAge)
		fmt.Printf("\tSecure=%t", curCk.Secure)
		fmt.Printf("\tHttpOnly=%t", curCk.HttpOnly)
		fmt.Printf("\tRaw=%s", curCk.Raw)
		fmt.Printf("\tUnparsed=%s", curCk.Unparsed)
	}
}

func TestCookie(t *testing.T) {
	var myCookieArray []*http.Cookie
	var myCookieJar *cookiejar.Jar

	myCookieArray = nil
	myCookieJar, _ = cookiejar.New(nil)
	//host:="http://api.xuechebu.com"
	//url:="/usercenter/userinfo/login?osversion=5.0.2&ossdk=21&passwordmd5=ec4ca7cbb7f22d8a85525db3787528da&imei=458606304AD7370C556B9E314D33CA5F&username=13661303427&appversion=2.0.0&version=2.0.0&ipaddress=172.19.190.2&os=an"
	host:="http://45.62.101.92:8888"
	url1:="/login"

	//GetUrlRespHtml(host+url,myCookieJar,myCookieArray)
	data:= url.Values{"username":{"jspdba"},"password":{"wuchaofei1"}}
	HttpPostForm(host+url1,&data,myCookieJar,myCookieArray)
	//PrintCurCookies(myCookieArray)
	//GetUrlRespHtml(host+url,myCookieJar,myCookieArray)
	//PrintCurCookies(myCookieArray)
}

func TestForTest(t *testing.T) {
	utils.ForTest1()
}
