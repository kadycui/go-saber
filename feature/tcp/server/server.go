package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {

	// 处理完关闭链接
	defer conn.Close()

	// 针对当前连接做发送和接受操作
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte

		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("读取链接失败,err:%v\n", err)
		}

		recv := string(buf[:n])
		fmt.Printf("收到的数据: %v\n", recv)

		// 将接收到的数据返回给客户端
		_, err = conn.Write([]byte("ok"))
		if err != nil {
			fmt.Printf("当前链接额写入失败, err: %v\n", err)
			break
		}
	}

}

func main() {
	// 建立tcp服务
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("监听失败, err:%v", err)
		return
	}

	for {
		// 等待客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("接收失败, err:%v", err)
			continue
		}

		// 启动一个单独的goroutine去处理连接
		go process(conn)

	}

}
