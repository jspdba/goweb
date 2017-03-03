package main

import (
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(client)
	}
}


func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()
	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("VER===%#x",b[0])
	log.Printf("NMETHODS===%#x",b[1])
	log.Printf("METHODS===%#x",b[2])
	if b[0] == 0x05 {//only process socket5
		//client response no auth
		client.Write([]byte{0x05, 0x00})
		n, err = client.Read(b[:])
		log.Printf("CMD===%#x",b[1])
		log.Printf("RSV===%#x",b[2])
		var host, port string
		switch b[3] {
		case 0x01:
			//IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //address
			host = string(b[5 : n-2]) //b[4] length of host
		case 0x04: //ip v6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		log.Printf("%b,%b",int(b[n-2]<<8),int(b[n-1]))
		port = strconv.Itoa(int(b[n-2]<<8) | int(b[n-1]))
		log.Println(host,port)
		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			log.Println(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) //success and send data redirect
		go io.Copy(server, client)
		io.Copy(client, server)
	}
}