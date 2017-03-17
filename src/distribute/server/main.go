package main

import (
	dis "distribute/context"
	"time"
)

//测试zk连接
func main() {
	/*
	连接并创建节点
	conn :=GetConnect([]string{"45.62.101.92:2181"})
	defer conn.Close()
	conn.Create("/go_servers",nil,0,zk.WorldACL(zk.PermAll))
	time.Sleep(2*time.Second)*/

	/*
	连接并创建节点
	zkList := []string{"45.62.101.92:2181"}
	conn := GetConnect(zkList)

	defer conn.Close()
	conn.Create("/go_temp", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)*/

	/*conn :=GetConnect([]string{"45.62.101.92:2181"})
	defer conn.Close()

	children, _, err := conn.Children("/go_servers")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v \n", children)*/
	startServerMain()
}

/*
func GetConnect(zkList []string) (conn *zk.Conn){
	conn,_,e:= zk.Connect(zkList,10*time.Second)
	if e!=nil{
		beego.Error(e)
		return
	}
	return
}*/


func startClientMain() {
	for i := 0; i < 100; i++ {
		dis.StartClient()
		time.Sleep(1 * time.Second)
	}
}
func startServerMain() {
	go dis.StarServer("127.0.0.1:8897")
	go dis.StarServer("127.0.0.1:8898")
	go dis.StarServer("127.0.0.1:8899")

	a := make(chan bool, 1)
	<-a
}