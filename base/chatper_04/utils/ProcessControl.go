package utils

import (
	"fmt"
	"runtime"
)

func Demo() {
	num := 101
	if num > 100 {
		fmt.Println(num, "> 100")

	} else if num == 100 {
		fmt.Println(num, " = 100")
	} else {
		fmt.Println(num, "< 100")
	}
}

func Demo1() {
	if num := runtime.NumCPU(); num > 1 {
		fmt.Println("程序使用的CPU核数为:", num)
	}
}

func Demo2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("i的值是:", i)
	}

}

func Demo3() {
	i := 1
	for {
		for {
			if i > 5 {
				fmt.Println("跳出内层for循环")
				break
			}
			fmt.Println(i)
			i++
		}
		fmt.Println("跳出外层循环")
		break

	}
}

func Demo4() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}

}

func Demo5(a int, b int) {
	switch a + b {
	case 1:
		fmt.Println("a+b = 1")
	case 2:
		fmt.Println("a+b = 2")
	case 3:
		fmt.Println("a+b = 3")
	case 4:
		fmt.Println("a+b = 4")

	}

}

func Demo6() {
	// 如果我们需要无条件强制执行后面的 case，可以使用fallthrough关键字
	switch {
	case true:
		fmt.Println("case 1为true")
		fallthrough
	case false:
		fmt.Println("case 2为false")
		fallthrough
	default:
		fmt.Println("默认 case")
	}

}

func Demo7() {
	fmt.Println("goto")
	goto sign
	fmt.Println("这是无效代码")
sign:
	fmt.Println("hello hello")
}

func Demo8(n int) {
	for i := 1; i <= n; i++ {
		// 打印＊前先打印空格，规律为总层数-当前层数
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		// //k表示每层打印多少＊，规律为 2 ＊ i - 1
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()

		/*
			       *
			      ***
			     *****
			    *******
			   *********
			  ***********
			 *************
			***************

		*/
	}
}
