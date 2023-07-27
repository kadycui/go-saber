package utils

import "fmt"

func LightCopy() {
	a := []int{1, 2, 3}
	b := a   // 浅拷贝
	b[0] = 4 // 修改b中的第一个元素

	fmt.Printf("a 1: %d\t内存地址: %p \n", a, a)
	fmt.Printf("b 2: %d\t内存地址: %p \n", b, b)

}
