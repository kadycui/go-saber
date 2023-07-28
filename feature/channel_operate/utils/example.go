package utils

import (
	"fmt"
	"time"
)

func writeData(ic chan int) {
	for i := 0; i < 50; i++ {
		// 数据放入管道
		ic <- i
		fmt.Println("写入数据 ", i)

	}

	close(ic)

}

func readData(ic chan int, ec chan bool) {
	for {
		v, ok := <-ic
		if !ok {
			break
		}

		fmt.Printf("读取数据=%v\n", v)
	}

	// 传入读取完成信号
	ec <- true
	close(ec)
}

func ReadWriteChan() {
	// 创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	// 开启两个 goroutine
	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)
	for {
		_, ok := <-exitChan
		if !ok {
			fmt.Println("主函数结束")
			break
		}
	}

}

// 存入管道
func putNum(ic chan int) {
	for i := 1; i < 8000; i++ {
		ic <- i

	}

	close(ic)
}

// 从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
func primeNum(ic chan int, pc chan int, ec chan bool) {

	var flag bool

	for {
		num, ok := <-ic
		if !ok {
			break
		}

		flag = true

		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}

		}

		if flag {
			pc <- num
		}

	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	// 这里我们还不能关闭 pc, 其他goroutine还在写入数据
	ec <- true

}

func FindPrimeNum() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)

	// 标识退出的管道
	exitChan := make(chan bool, 4)

	start := time.Now().Unix()

	// 开启一个goroutine, 存数据
	go putNum(intChan)

	//开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就
	//放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)

	}

	// 主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan

		}

		end := time.Now().Unix()

		fmt.Println("使用goroutine耗时: ", end-start)

		close(primeChan)
	}()

	// 遍历primeChan  取出素数
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		// 输出结果
		fmt.Printf("素数=%d \n", res)

	}

	fmt.Println("主线程结束!!!")

}
