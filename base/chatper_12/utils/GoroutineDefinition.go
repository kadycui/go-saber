package utils

import (
	"fmt"
	"runtime"
	"time"
)

// goroutine
/*
在Go语言中，每一个并发的执行单元叫作一个goroutine。想要编写一个并发任务，只需要在调
用的函数前面添加go关键字，就能使这个函数以协程的方式运行。

go 函数名(函数参数)
*/

func Task1() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Task1的任务!")
		time.Sleep(time.Second * 3)
	}
}

func Task2() {
	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Task2的任务!")
		time.Sleep(time.Second * 1)
	}
}

func Demo() {
	go Task1()
	go Task2()

	for {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Demo的任务!")
		time.Sleep(time.Second * 2)

	}
}

func Demo1() {
	/*
		运行后将不会有任何输出，这是因为运行go Task1()，程序会立刻返回到main函
		数，而main函数后面没有任何代码逻辑，程序就会判断为执行完毕，终止所有协程。因此，想要
		让Task1函数执行，可以在main函数添加一些等待逻辑，例如Sleep()
	*/
	go Task1()

	time.Sleep(time.Second * 100)
}

// 使用匿名函数创建 goroutine

/*
go func(参数列表) {
 函数体
} (调用参数列表)
*/

func Demo2() {
	go func() {
		for {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "正在处理Task的任务!!")
			time.Sleep(time.Second * 3)
		}
	}()
	time.Sleep(time.Second * 100)
}

// runtime包
/*
Gosched()、Goexit()、GOMAXPROCS()。


Gosched()使当前Go协程放弃处理器，以让其他Go协程运行。它不会挂起当前Go协程，因此当前Go协程未来会恢复执行。
Go语言的协程是抢占式调度的，当遇到长时间执行或者进行系统调用时，会主动把当前goroutine的CPU (P)转让出去，
让其他goroutine能被调度并执行。一般出现如下几种情况，goroutine就会发生调度：
◇ syscall。
◇ C函数调用（本质上和syscall一样）。
◇ 主动调用runtime.Gosched。
◇ 某个goroutine的调用时间超过100 ms，并且这个goroutine调用了非内联的函数。
*/

func Demo3() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("main")
	}
}

func Demo4() {
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {

		// 使用runtime.Gosched()来阻止其获取控制权
		runtime.Gosched()
		fmt.Println("main")
	}
}

func Task11() {
	defer fmt.Println("task11 stop!!")
	fmt.Println("task11 start!!")
	fmt.Println("task11 work!!")

}

func Task12() {
	defer fmt.Println("task12 stop!!")
	fmt.Println("task12 start!!")

	// Goexit()效果和return一样
	runtime.Goexit()
	fmt.Println("task12 work!!")

}

func Demo5() {
	go Task11()
	go Task12()

	time.Sleep(time.Second * 3)
}

// GOMAXPROCS(n int)函数可以设置程序在运行中所使用的CPU数，在以后的编程中是用得最多的。Go语言程序默认会使用最大CPU数进行计算
/*
func GOMAXPROCS(n int) int

GOMAXPROCS()设置可同时执行的最大CPU数，并返回先前的设置。
若n ＜ 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过NumCPU查询
*/

func Demo6() {

	cn := runtime.NumCPU()
	n := runtime.GOMAXPROCS(2)
	fmt.Println("当前设备的逻辑cpu数:", cn)
	fmt.Println("先前的cpu核数设置为:", n)

	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		go func() {
			// 耗时任务
			a := 999999 ^ 9999999
			a = a + 1
		}()
	}
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime))

}
