package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Demo1() {
	num := 1
	var p *int
	p = &num

	// 一般情况下，我们将指针变量的类型声明为*int，变量名为“p”开头（指代“point”）的单词，如“p”或“ptr”
	fmt.Println("num变量的地址为:", p)

	// 使用操作符“&”取变量地址，取得的地址值可以赋给指针变量
	fmt.Println("指针变量p的地址为:", &p)

	// 通过在指针变量前面加上“*”符号可以获取指针所指向地址值的内容
	fmt.Println("指针变量p所指的内容:", *p)

	// 由于指针变量本身也是变量，因此指针变量在计算机内存中也有自己的地址
	/*
	 * num变量的地址为: 0xc0000180b8
	 * 指针变量p的地址为: 0xc000006030
	 */
}

func Demo2() {

	var p *int
	fmt.Println("指针变量指向的地址为:", p)
	// p指针声明后其值为nil，这时如果获取指针p指向的地址内容，则会出错
	// fmt.Println("指针变量p所指向的内容:", *p)
}

func Demo3() {
	num := 1
	var p *int
	p = &num
	fmt.Println("指针变量p所指向的内容:", *p, ",p的内存地址是", &p)

	*p = 10
	fmt.Println("指针变量p所指向的内容:", *p, ",p的内存地址是", &p)

	var pp *int
	pp = new(int)
	fmt.Println("指针变量pp所指向的内容:", *pp, ",pp的内存地址是", &pp)
	*pp = 10
	fmt.Println("指针变量pp所指向的内容:", *pp, ",pp的内存地址是", &pp)

}

func Demo4() {

	// + 拼接字符串
	//a := "12345678"
	//b := "910"
	//c := a + b
	//fmt.Println(c)

	// 字节缓冲
	a := "1111111"
	b := "2222222"

	var c bytes.Buffer
	c.WriteString(a)
	c.WriteString(b)
	fmt.Println(c.String())
	fmt.Println(reflect.TypeOf(c))

}

func Demo5() {
	// 字符串截取
	str := "举头望明月"
	index := strings.Index(str, "明")
	fmt.Println(index)
	fmt.Println(str[index:])

	// 只取最后出现的字符
	str2 := "Go语言,Python语言"
	index2 := strings.LastIndex(str2, "语")
	fmt.Println(index2)
	fmt.Println(str2[index2:])

	fmt.Println("--------------------------------------")

	// go语言无法对数组进行修改,只能将字符串转换为字节数组后再进行操作
	str3 := "Go语言"
	bs := []byte(str3)
	fmt.Println(bs, reflect.TypeOf(bs))

	for i := 0; i < 2; i++ {
		bs[i] = ' '
	}
	fmt.Println(string(bs))

}

func Demo6() {
	var day = 1
	var hour = 24
	str := fmt.Sprintf("%d天包含%d个小时", day, hour)
	fmt.Println(str)
}
