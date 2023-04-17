package utils

/*
1.应用层
这一层是最靠近用户的OSI层，为用户的应用程序（例如电子邮件、文件传输和终端仿真）提
供网络服务。
常见的协议有：HTTP、FTP、TFTP、SMTP、SNMP、DNS、TELNET、HTTPS、POP3、DHCP。

2.表示层
表示层可确保一个系统的应用层所发送的信息能够被另一个系统的应用层读取。例如，PC程
序与另一台计算机进行通信，其中一台计算机使用扩展二进制编码的十进制交换码（EBCDIC），
而另一台则使用美国信息交换标准码（ASCII）来表示相同的字符。如有必要，表示层会通过使用
一种通用格式来实现多种数据格式之间的转换。
常见的协议有：JPEG、ASCII、DECOIC。

3.会话层
会话层通过传输层（端口号：传输端口与接收端口）建立数据传输的通路。它主要在系
统之间发起会话或者接受会话请求（设备之间需要互相认识可以是IP，也可以是MAC或者是主
机名）。
常见的协议有：RPC、SCP、SSH、ZIP。

4.传输层
传输层定义了一些传输数据的协议和端口号（WWW端口80等），如：TCP（传输控制协议，
传输效率低，可靠性强，用于传输可靠性要求高、数据量大的数据），UDP（用户数据报协议，与
TCP特性恰恰相反，用于传输可靠性要求不高、数据量小的数据，如QQ聊天数据就是通过这种方
式传输的）。传输层主要是将从下层接收的数据进行分段和传输，到达目的地址后再进行重组。
我们常常把这一层数据叫作段。
常见的协议有：TCP、UDP。

5.网络层
网络层为位于不同地理位置的网络中的两个主机系统之间提供连接和路径选择。互联网的发
展使从世界各站点访问信息的用户数大大增加，而网络层正是管理这种连接的层。
常见的协议有：ICMP、IGMP、IP、RARP。

6.数据链路层
数据链路层定义了如何格式化数据以进行传输，以及如何控制对物理介质的访问。这一层通
常还提供错误检测和纠正，以确保数据的可靠传输。常见的协议有：PPP、IEEE 802.3/802.2。

7.物理层
物理层主要定义物理设备标准，如网线的接口类型、光纤的接口类型、各种传输介质的传输
速率等。它的主要作用是传输比特流（就是由1、0转化为电流强弱来进行传输，到达目的地后再
转化为1、0，也就是我们常说的数模转换与模数转换）。这一层的数据叫作比特。
常见的协议有：Ethernet 802.3

*/

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Demo1() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	log.Println("Connect Success")
}

func Demo2() {
	// 尝试连接本地3306接口
	conn, err := net.Dial("tcp", ":13306")
	if err != nil {
		log.Fatal("Connect Fail!", err)
	}
	defer conn.Close()
	log.Println("Connect Success!")
}

func Demo3() {
	// 设置超时时间 3 秒
	conn, err := net.DialTimeout("tcp", "www.baidu.com:81", time.Second*3)
	if err != nil {
		log.Fatal("连接失败", err)
	}
	defer conn.Close()
	log.Println("连接成功")
}

func Demo4() {
	//Windows环境:  nc -l -p 1234
	// 尝试连接本地3306接口
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("Connect Fail!", err)
	}
	defer conn.Close()
	log.Println("Connect Success!")

	// 发送数据
	conn.Write([]byte("Test!!\n"))

	// 接收数据
	var buf = make([]byte, 10)
	conn.Read(buf)
	log.Println(buf)
}

func HttpDemo5() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal("Connect Fail!", err)
	}
	defer conn.Close()
	log.Println("Connect Success!")

	// 发送Http内容
	conn.Write([]byte("GET / HTTP/1.1\\r\\nHost: www.baidu.com\\r\\\n nUser-Agent: curl/7.55.1\\r\\nAccept: ＊/＊\\r\\n\\r\\n"))

	var buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println(string(buf))

}
