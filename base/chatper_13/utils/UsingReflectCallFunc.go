package utils

import (
	"fmt"
	"reflect"
)

/*使用反射调用函数需要用到reflect.ValueOf()方法传入想要反射的函数名，
获取到reflect.Value对象，再通过reflect.Value对象的Call方法调用该函数

func (v Value) Call(in []Value) []Value
*/

func Equal(a, b int) bool {
	if a == b {
		return true
	}
	return false
}

func Demo4() {
	// 反射调用函数需使用 ValueOf
	valueOfFunc := reflect.ValueOf(Equal)

	// 构造函数参数
	args := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)}

	// 通过反射调用函数计算
	result := valueOfFunc.Call(args)

	fmt.Println("函数运行结果: ", result[0].Bool())
}
