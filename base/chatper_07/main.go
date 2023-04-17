package main

import (
	"fmt"
	"gostart/base/chatper_07/utils"
)

//func addSub(x int, y int) (sum int, sub int) {
//	sum = x + y
//	sub = x - y
//	return sum, sub
//}

func addall(slice ...int) {
	sum := 0
	for _, value := range slice {
		sum = sum + value
	}

	fmt.Println(sum)
}

func add(num ...int) {
	addall(num...)

}

func addOne(i int) func() int {
	return func() int {
		i++
		return i
	}

}

func main() {
	//a := 1
	//b := 2
	//
	//// 函数赋值
	//f1 := addSub
	//sum, sub := f1(a, b)
	//fmt.Println(a, "+", b, "=", sum)
	//fmt.Println(a, "-", b, "=", sub)

	fmt.Println("1+2+...+9+10=", utils.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// 匿名函数
	func(data string) {
		fmt.Println("hello " + data)
	}("Go!")

	f1 := func(data string) {
		fmt.Println("Hello " + data)
	}
	f1("world!")

	utils.ClosureFunc()
	fmt.Println("==========================")

	/*
		addOne函数返回了一个闭包函数，通过定义a1和a2变量，创建了两个闭包的实例（引用环境
		不同导致）。
		每次调用闭包实例，i的值都会在原有的基础上加1。从打印的结果可以看到，两个闭包实例的
		地址完全不同，两个闭包的调用结果互不影响。
	*/
	//a1 := addOne(0)
	//fmt.Println(a1()) //0+1=1
	//fmt.Println(a1()) //1+1=2
	//a2 := addOne(10)
	//fmt.Println(a2())
	//fmt.Print("a1闭包的地址为：")
	//fmt.Printf("%p\n", &a1)
	//fmt.Print("a2闭包的地址为：")
	//fmt.Printf("%p\n", &a2)

	//utils.Demo2()
	//fmt.Println("==========================")
	//utils.TcpSend()
	//
	//utils.Demo3()
	//utils.Demo4()
	utils.Demo5()

}
