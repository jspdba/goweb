package service

import (
	"web/models"
	"net/http"
	"net/url"
	"time"
	"github.com/astaxie/beego"
	"net"
)

func CheckProxyValid(proxyIp *models.ProxyIp, timeout time.Duration)  bool{
	/*conn, err := net.DialTimeout("tcp", proxyIp.Ip+":"+strconv.Itoa(proxyIp.Port),2*time.Second)
	if err!=nil{
		return false
	}
	defer conn.Close()
	return true*/
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://"+proxyIp.Ip+":"+proxyIp.Port)
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
			Dial: (&net.Dialer{
				Timeout:   3*time.Second,
				Deadline:  time.Now().Add(timeout),
				KeepAlive: 3*time.Second,
			}).Dial,
		},
		Timeout: timeout,
	}
	_, err := client.Get("http://www.ip138.com")
	if err != nil {
		beego.Info(err)
		return false
	}
	return true
}
