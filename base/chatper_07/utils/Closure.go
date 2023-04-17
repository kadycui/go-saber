package utils

import "fmt"

// 闭包函数

// 匿名函数由于在函数体内部引用了外部的自由变量num而形成了闭包。闭包每次对num变量的加1操作都是对变量num引用的修改

func ClosureFunc() {
	num := 1
	fmt.Printf("%p\n", &num)
	func() {
		num++
		fmt.Println(num)
		fmt.Printf("%p\n", &num)
	}()

	func() {
		num++
		fmt.Println(num)
		fmt.Printf("%p\n", &num)
	}()
}
