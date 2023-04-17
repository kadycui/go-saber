package utils

import (
	"log"
	"net"
)

func UdpDemo1() {
	// 创建地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:8080")
	checkErr(err)
	// 建立连接
	conn, err := net.DialUDP("udp", nil, udpaddr)
	checkErr(err)

	defer conn.Close()
	log.Println("连接成功!!")

}

func UdpDemo2() {
	// 创建地址
	udpaddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:1234")
	checkErr(err)
	// 建立连接
	conn, err := net.DialUDP("udp", nil, udpaddr)
	checkErr(err)

	defer conn.Close()
	log.Println("连接成功!!")

	// 发送数据
	conn.Write([]byte("Hello\r\n"))

	// 接收数据
	var buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println("接收的数据", string(buf))

}
