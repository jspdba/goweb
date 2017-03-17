package context

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"fmt"
)

var (
	Hosts =[]string{"45.62.101.92:2181"}
)

func GetConnect() (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect(Hosts, 10 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func RegistServer(conn *zk.Conn, host string) (err error) {
	_, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetServerList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/go_servers")
	return
}
