package utils

import "fmt"

func Log(name string, i interface{}) {
	fmt.Printf("Name = %s, Type = %T, value = %v\n", name, i, i)
}

/*
空接口（interface{}）是Go语言中最特殊的接口。在Java语言中，所有的类都继承自一个基类
Object，而Go中的interface{}接口就相当于Java语言里的Object。
在Go语言中，空接口不包含任何方法，也正因如此，所有的类型都实现了空接口，因此空接
口可以存储任意类型的数值。
*/

func Demo7() {
	var v1 interface{} = 1
	var v2 interface{} = &v1
	var v3 interface{} = true
	var v4 interface{} = "abc"
	var v5 interface{} = struct {
		Name string
	}{"小明"}
	var v6 interface{} = &v5

	Log("v1", v1)
	Log("v2", v2)
	Log("v3", v3)
	Log("v4", v4)
	Log("v5", v5)
	Log("v6", v6)

}

// 空接口取值
func Log2(args ...interface{}){
	for num, arg := range args{
		fmt.Printf("Index => %d, Value => %v\n",num,arg)
	}

}

func Demo8(){
	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "abc"
	s[2] = struct{
		Num int
	}{8}

	// 可变长参数
	fmt.Println("------将切片拆散------")
	Log2(s ...)
	fmt.Println("====直接传入切片====")
	Log2(s)

}
