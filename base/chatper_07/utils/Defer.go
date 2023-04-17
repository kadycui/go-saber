package utils

import (
	"fmt"
	"net"
)

func Demo2() {
	fmt.Println("start now!!")
	defer fmt.Println("这是第一句defer语句")
	defer fmt.Println("这是第二句defer语句")
	defer fmt.Println("这是第三句defer语句")
	fmt.Println("end!!")

}

func TcpSend() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err == nil {
		defer conn.Close()
		fmt.Println("remote address:", conn.RemoteAddr())
	}
	fmt.Println("error:", err)
}
