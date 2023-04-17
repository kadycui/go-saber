package utils

import (
	"log"
	"net"
	"time"
)

func TcpDemo1() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("服务启动失败", err)
	}

	defer l.Close()
	log.Println("服务启动成功!")

	//阻塞等待连接
	c, err := l.Accept()

	// 设置连接超时
	c.SetDeadline(time.Now().Add(time.Second))

	// 设置读取超时
	c.SetReadDeadline(time.Now().Add(time.Second))

	// 设置写入超时
	c.SetWriteDeadline(time.Now().Add(time.Second))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TcpDemo2() {
	l, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer l.Close()
	log.Println("服务启动成功!")

	// nc -nv 127.0.0.1 8080
	//阻塞等待连接
	c, err := l.Accept()
	checkErr(err)

	// 设置连接超时
	c.SetDeadline(time.Now().Add(time.Second * 20))

	// 设置读取超时
	c.SetReadDeadline(time.Now().Add(time.Second * 20))

	// 设置写入超时
	c.SetWriteDeadline(time.Now().Add(time.Second * 20))

	defer c.Close()

	// 接受信息
	var buf = make([]byte, 10)
	log.Println("开始连接!!")
	n, err := c.Read(buf)
	checkErr(err)
	log.Println("接收字节数：", n, "接收字内容为：", string(buf))

}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		var buf = make([]byte, 10)
		log.Println("开始读取连接!!!")
		n, err := c.Read(buf)
		checkErr(err)

		log.Printf("读取 %d 个字节, 内容是 %s \n", n, string(buf[:n]))

	}

}

func TcpDemo3() {
	l, err := net.Listen("tcp", ":8080")
	checkErr(err)
	defer l.Close()
	log.Println("服务启动成功!")

	for {
		c, err := l.Accept()
		checkErr(err)

		// 开启协程
		go handleConn(c)
	}

}
