package utils

/*
指的是这样一个数列：1、1、2、3、5、8、13、21、34……这个数列从第三项开始，每一项都等于前两项之和
*/

func F1(n int) (res int) {
	a := 1
	b := 1
	for i := 2; i < n; i++ {
		c := b
		b = a + b
		a = c

	}
	return b

}

func F2(n int) (res int) {
	// 递归方式
	if n == 1 || n == 2 {
		return 1
	} else {
		return F2(n-2) + F2(n-1)
	}
}

func F3(n int) (res int) {
	// n的阶乘
	if n == 1 {
		return 1
	} else {
		res := n * F3(n-1)
		return res
	}
}
