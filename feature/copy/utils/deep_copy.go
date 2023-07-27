package utils

import "fmt"

// 浅拷贝
func DeepCopy() {
	a := []int{1, 2, 3}
	b := make([]int, len(a)) // 创建一个新的切片

	copy(b, a) // 复制a中的所有元素到b中
	// copy函数不会扩容，也就是要复制的 slice 比原 slice 要大的时候，会移除多余的

	b[0] = 4 // 修改b中的元素

	fmt.Printf("a 1: %d\t内存地址: %p \n", a, a)
	fmt.Printf("b 2: %d\t内存地址: %p \n", b, b)

}
