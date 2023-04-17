package utils

import (
	"log"
	"net"
)

func UdpServerDemo1() {
	// 创建一个UDP地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	checkErr(err)

	// 创建UDP服务
	conn, err := net.ListenUDP("udp", udpaddr)
	checkErr(err)

	defer conn.Close()
	log.Println("UDP服务创建成功!!")

	var buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println("接受的数据", string(buf))

	_, responseAddr, err := conn.ReadFromUDP(buf)
	checkErr(err)

	conn.Write([]byte("Hello Write \r\n"))
	conn.WriteToUDP([]byte("Hello WriteToUdp \r\n"), responseAddr)
}
