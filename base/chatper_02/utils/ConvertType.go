package utils

import (
	"fmt"
	"reflect"
)

func ConvertType() {
	var a int32 = 1234567891
	fmt.Println("变量a的值为:", a, ", 变量类型为: ", reflect.TypeOf(a))
	fmt.Println("转换变量a的类型为int16, 变量a的值变为:", int16(a), ", 变量类型为: ", reflect.TypeOf(int16(a)))
}
