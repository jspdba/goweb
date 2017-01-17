package test

import (
	"testing"
	"flag"
	"net/http"
	"os"
	"io"
	"net"
	"log"
	"web/utils"
)

var get_ip = flag.String("get_ip", "", "external|internal")
func get_external() {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	os.Exit(0)
}
func get_internal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				os.Stdout.WriteString(ipnet.IP.String() + "\n")
			}
		}
	}
	os.Exit(0)
}
func Test_getIp(t *testing.T) {
	log.Println(utils.GetIp())
}
