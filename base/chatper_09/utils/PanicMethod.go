package utils

import "fmt"

// 一般而言，只有当程序发生不可逆的错误时，才会使用panic方法来触发宕机。panic方法是
//Go语言的一个内置函数，使用panic方法后，程序的执行将直接中断。

/*
func panic(v interface{})
*/

func Demo4() {
	// 宕机
	panic("Serious bug")
	fmt.Println("Invalid code") // 程序退出, 无法执行改行代码
}

// 宕机恢复

func protect() {
	defer func() {
		fmt.Println("func protect exit")
	}()
	panic("Serious bug")
}

func Demo5() {
	defer func() {
		fmt.Println("func main exit")
	}()
	protect()
	fmt.Println("Invalid code")

}

/*
1.protect函数内的panic方法触发宕机。
2.由于protect函数内的匿名函数通过defer语句延迟执行，在panic方法触发宕机后，且在退出
protect函数前，会执行protect函数中的匿名函数，打印“func protect exit”。
3.由于main函数内的匿名函数通过defer语句延迟执行，在main函数退出前会执行main函数中
的匿名函数，打印“func main exit”。
4.程序退出。
*/

func protect2() {
	defer func() {
		if err := recover(); err != nil { //recover()获取panic()传入的参数
			fmt.Println(err)
		}
	}()
	panic("Serious bug")
}

func Demo6() {
	protect2()
	fmt.Println("valid code")

}
