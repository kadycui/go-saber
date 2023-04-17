package utils

import (
	"fmt"
	"reflect"
)

// Go语言使用reflect.TypeOf来获取类型信息，使用reflect.ValueOf来获取变量值的信息
// v.Kind() 类型名
// v.Elem() 指向的变量

func checkValue(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Int {
		// 方法一
		var v1 int = int(v.Int())

		// 方法二
		var v2 int = v.Interface().(int)

		fmt.Println(v1, v2)
	}

}

func Demo3() {
	var num int = 10
	valueOfNum := reflect.ValueOf(num)
	fmt.Println("valueOfNum")
	checkValue(valueOfNum)

	valueOfNumPtr := reflect.ValueOf(&num)
	fmt.Println("valueOfNumPtr")
	checkValue(valueOfNumPtr)

}
