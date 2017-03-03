package test

import (
	"testing"
	"flag"
	"strings"
	"fmt"
	"log"
)

func TestFlag(t *testing.T) {
	var ip string
	var list string
	flag.StringVar(&ip, "l", ":9897", "-l=0.0.0.0:9897 指定服务监听的端口")
	flag.StringVar(&list, "d", "127.0.0.1:1789,127.0.0.1:1788", "-d=127.0.0.1:1789,127.0.0.1:1788 指定后端的IP和端口,多个用','隔开")
	flag.Parse()
	trueList := strings.Split(list, ",")
	if len(trueList) <= 0 {
		fmt.Println("后端IP和端口不能空,或者无效")
	}
	log.Println(ip,list,trueList)
}
