package utils

import "fmt"

func Demo1() {
	const a = 3.1415
	const b = "hello hello"
	fmt.Println(a)

}

func Demo2() {
	// 定义常量组
	const (
		a = 3.1415
		b
		c
		d = 100
		e = "hello"
	)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}

func Demo3() {
	// 常量枚举
	const (
		a = iota          // 0
		b                 // iota += 1    iota=1
		c = "hello hello" // iota += 1    iota=2
		d                 // iota += 1    iota=3
		e = iota          // iota += 1    iota=4
	)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
