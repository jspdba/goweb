package utils

import (
	"net/http"
	"io/ioutil"
	"github.com/astaxie/beego"
	"strings"
	"net/url"
	"encoding/base64"
	"net"
	"time"
	"fmt"
	"net/http/cookiejar"
)
type user struct {
	userid string
	cookieStore string
}

func login(url string){

}

func logout(url string){

}

func httpGet(url string) (content string){
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
	}
	resp, err := client.Get(url)
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
		content=string(body)
	}
	return
}

func httpPost(url string) (content string){

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
	}

	resp, err := client.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
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
		content=string(body)
	}

	return
}
func httpPostForm(url string, data url.Values) (content string){
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
	}
	resp, err := client.PostForm(url,data)
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
	return

}

func httpDoPost(url,cookies string) (content string){
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
	}

	req, err := http.NewRequest("POST", url, strings.NewReader("name=cjb"))
	if err != nil {
		beego.Error(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookies)

	resp, err := client.Do(req)

	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
			return
		}
		content=string(body)
	}
	return
}

//用base64进行编码
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))

}

//用base64进行解码
func base64Decode(src []byte) ([]byte, error) {

	return base64.StdEncoding.DecodeString(string(src))

}


func ForTest() {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "M_WEIBOCN_PARAMS",
		Value:  "rl%3D1",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "SUB",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "_T_WM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "gsid_CTandWM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	u, _ := url.Parse("http://weibo.cn/search/?vt=4")
	//jar.SetCookies(u, cookies)
	fmt.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
	}
	postData := url.Values{}
	postData.Set("keyword", "尹相杰")
	postData.Set("smblog", "搜微博")
	req, _ := http.NewRequest("POST", "http://weibo.cn/search/?vt=4", strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	_ = body
	fmt.Println(jar.Cookies(u))
	//fmt.Println(string(body))
}

func myRedirect(req *http.Request, via []*http.Request) (e error) {
	beego.Info(req.URL.String())
	beego.Info(via[0].URL.String())
	return http.ErrUseLastResponse
	//return
}

func ForTest1() {
	jar, _ := cookiejar.New(nil)
	/*var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "M_WEIBOCN_PARAMS",
		Value:  "rl%3D1",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "SUB",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "_T_WM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)
	cookie = &http.Cookie{
		Name:   "gsid_CTandWM",
		Value:  "xxx",
		Path:   "/",
		Domain: ".weibo.cn",
	}
	cookies = append(cookies, cookie)*/
	u, _ := url.Parse("http://45.62.101.92:8888/login")
	//jar.SetCookies(u, cookies)
	//fmt.Println(jar.Cookies(u))
	client := &http.Client{
		Jar: jar,
		CheckRedirect: myRedirect,
	}
	postData := url.Values{}
	postData.Set("username", "jspdba")
	postData.Set("password", "wuchaofei1")
	req, _ := http.NewRequest("POST", "http://45.62.101.92:8888/login", strings.NewReader(postData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(nil)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	_ = body

	fmt.Println(resp.Cookies())

	jar.SetCookies(u,resp.Cookies())
	fmt.Println(jar.Cookies(u))
	fmt.Println(*jar)
	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
}