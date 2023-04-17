package utils

import "fmt"

func Demo4() {

	// 算术运算符
	a := 1
	b := 2
	fmt.Println("a+b=", a+b)
	fmt.Println("a-b=", a-b)
	fmt.Println("a＊b=", a*b)
	fmt.Println("a/b=", a/b)
	fmt.Println("a%b=", a%b)
	a++
	fmt.Println("a++后a值为:", a)
	b--
	fmt.Println("b--后b值为:", b)
}

func Demo5() {
	// 比较运算符
	a := 1
	b := 2
	fmt.Println("a==b:", a == b)
	fmt.Println("a!=b:", a != b)
	fmt.Println("a>b:", a > b)
	fmt.Println("a<b:", a < b)
	fmt.Println("a>=b:", a >= b)
	fmt.Println("a<=b:", a <= b)
}

func Demo6() {
	// 赋值运算符
	var a = 10
	fmt.Println("a=", a)
	a += 2
	fmt.Println("a += 2,a=", a)
	a -= 2
	fmt.Println("a -= 2,a=", a)
	a *= 2
	fmt.Println("a ＊= 2,a=", a)
	a /= 2
	fmt.Println("a /= 2,a=", a)
	a %= 2
	fmt.Println("a %= 2,a=", a)
}

func Demo7() {
	// 位运算符
	a := 9
	b := 13
	fmt.Println("a&b = ", a&b)
	fmt.Println("a|b = ", a|b)
	fmt.Println("a^b = ", a^b)
	fmt.Println("a<<2 = ", a<<2)
	fmt.Println("b>>2 = ", b>>2)

	/*
		1.按位与：将参与运算的两个数按二进制位展开后进行与运算，9按二进制位展开为1001，
		13按二进制位展开为1101，1001和1101的与运算结果为1001，即十进制9。
		2.按位或：将参与运算的两个数按二进制位展开后进行或运算，9按二进制位展开为1001，
		13按二进制位展开为1101，1001和1101的或运算结果为1101，即十进制13。
		3.按位异或：将参与运算的两个数按二进制位展开后进行异或运算，9按二进制位展开为
		1001，13按二进制位展开为1101，1001和1101的异或运算结果为0100，即十进制4。
		4.按位左移：将参与运算的数按二进制位展开后全部左移指定位数，9按二进制位展开为
		1001，全部左移2位后变为100100，即十进制36。
		5.按位右移：将参与运算的数按二进制位展开后全部右移指定位数，13按二进制位展开为
		1101，全部右移2位后变为11，即十进制3
	*/
}

func Demo8() {
	// 逻辑运算符
	a := true
	b := false
	fmt.Println("a&&b = ", a && b)
	fmt.Println("a||b = ", a || b)
	fmt.Println("!a = ", !a)

	d := 1
	var p *int
	p = &d
	/*
		正数1按二进制位展开为001，其中最左位0表示正数，取反操作对所有二进制位取反，结果为
		110，其中最左位1表示负数，即十进制-2
	*/
	fmt.Println("^a = ", ^d)
	fmt.Println("d的变量地址:", p)
}
