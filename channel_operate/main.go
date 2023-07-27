package main

import (
	"fmt"
)

// func recv(c chan int) {
// 	ret := <-c
// 	fmt.Println("接收成功", ret)

// }

func cter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

// 单向发送 out 通道， 单向接收 in 通道
func sqer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// 单向接收 in 通道
func prter(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
func main() {

	// 无缓冲通道(阻塞通道, 同步通道)
	// ch := make(chan int)

	// 有缓冲通道, 创建一个容量为2的通道(非阻塞)
	// ch := make(chan int, 1)

	// go recv(ch) // 启用 goroutine从通道接收值
	// ch <- 10
	// time.Sleep(time.Second)

	// fmt.Println("发送成功")

	out := make(chan int)
	in := make(chan int)

	go cter(out)
	go sqer(out, in)

	prter(in)

}
