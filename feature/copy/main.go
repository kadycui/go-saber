package main

import (
	"fmt"

	"github.com/kadycui/go-saber/feature/copy/utils"
)

/*
Go语言中有丰富的数据类型，除了基本的整型、浮点型、布尔型、字符串外，
还有数组、切片、结构体、函数、map、通道（channel）等。
Go 语言的基本类型和其他语言大同小异。

值类型：
	变量直接存储值，内存通常在栈中分配
	（属于值类型的数据类型有：int、float、bool、string、数组以及struct）。

引用类型：
	变量存储的是一个地址，这个地址存储最终的值，内存通常在堆中分配，通过GC回收
	（属于引用类型的的数据类型有：指针、slice、map、chan等）。
*/

func main() {

	utils.DeepCopy()
	fmt.Println("--------------------------------------")
	utils.LightCopy()

}

/*
归根结底，使用 := 进行拷贝时：

1. 会产生深拷贝还是浅拷贝取决于被拷贝数据的数据类型：
	如果是值类型，就会产生深拷贝
	如果是引用类型，就会产生浅拷贝

2. 深浅拷贝的区别：
	深拷贝：光拷贝值，地址不相关。拷贝结束两变量互不影响。
	浅拷贝：拷贝地址。两变量指向同一地址。拷贝结束也相关，改变一个，另一个也跟着变。

3. 另外，对于引用类型，想实现深拷贝，就不能直接 := ，而是要先 new ，再赋值。
*/
