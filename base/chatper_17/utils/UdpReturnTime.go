package utils

import (
	"log"
	"net"
	"time"
)

func handelClient(conn *net.UDPConn) {

	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	checkErr(err)

	log.Println(n, remoteAddr)

	b := make([]byte, 1024)
	b = []byte(string(time.Now().String()))
	conn.WriteToUDP(b, remoteAddr)

}

func ReturnTime() {
	// 创建一个UDP地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	checkErr(err)

	// 创建UDP服务
	conn, err := net.ListenUDP("udp", udpaddr)
	checkErr(err)

	defer conn.Close()
	log.Println("UDP服务创建成功!!")
	for {
		handelClient(conn)
	}
}
