package main

import (
	"fmt"
	"reflect"
)

// ReturnData 匿名变量
func ReturnData() (int, int) {

	return 30, 40
}

func main() {

	fmt.Println("----------变量定义")
	var num int = 1
	var (
		a int    = 2
		b string = "nishi"
		c bool   = true
	)

	var name string
	// 多个短变量声明并初始化中, 应该至少有一个新的变量出现在左值中
	name, age := "Chuan", 50

	fmt.Println(num)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("name:", name)
	fmt.Println("age:", age)

	// 变量交换
	aa := 34
	bb := 78

	aa, bb = bb, aa
	fmt.Println("aa的值是", aa)
	fmt.Println("bb的值是", bb)

	fmt.Println("----------匿名变量")
	ab, _ := ReturnData()
	_, ac := ReturnData()
	fmt.Println(ab, ac)

	fmt.Println("----------数据类型")

	ad := 3
	ae := 2
	fmt.Println(ad / ae)

	af := 3.0
	ag := 2.0
	fmt.Println(af / ag)
	fmt.Println("变量af的类型是:", reflect.TypeOf(af))
	fmt.Println("变量ag的类型是:", reflect.TypeOf(ag))

	fmt.Println("----------字符串")
	s1 := "这是一个字符串!"
	s2 := `这是第一行
	这是第二行
	这是第三行
	`
	fmt.Println(s1)
	fmt.Println(s2)

	english := 'i'
	chinese := '我'
	fmt.Println(english)
	fmt.Println(chinese)

	fmt.Println("----------数据类型转换")
	var ba int16 = 97
	fmt.Println("变量ba的值为:", ba, "变量类型是:", reflect.TypeOf(ba))
	bc := int32(ba)
	fmt.Println("变量bb的值为:", bc, "变量类型是:", reflect.TypeOf(bc))
	fmt.Println("转换变量bc的类型为string", string(bc))

	var str22 = fmt.Sprintf("我叫 %s, 今年%s, 我住在%s", "张三", "8", "广州")
	fmt.Println(str22)

	fmt.Println("----------浮点数处理")

	var f1 = 7.0635
	var f2 float64 = 200

	var f3 = f1 * f2

	fmt.Printf("%.4f\n", f3)

}
