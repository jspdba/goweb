package context

import (
	"time"
	"fmt"
	"net"
	"io/ioutil"
	"errors"
	"math/rand"
)

func CheckError(err error){
	if err!=nil{
		panic(err)
	}
}

func StartClient() {
	// service := "127.0.0.1:8899"
	//获取地址
	serverHost, err := GetServerHost()
	if err != nil {
		fmt.Printf("get server host fail: %s \n", err)
		return
	}

	fmt.Println("connect host: " + serverHost)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
	CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	CheckError(err)

	result, err := ioutil.ReadAll(conn)
	CheckError(err)
	fmt.Println(string(result))

	return
}

func GetServerHost() (host string, err error) {
	conn, err := GetConnect()
	if err != nil {
		fmt.Printf(" connect zk error: %s \n ", err)
		return
	}
	defer conn.Close()
	serverList, err := GetServerList(conn)
	if err != nil {
		fmt.Printf(" get server list error: %s \n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		err = errors.New("server list is empty \n")
		return
	}

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}