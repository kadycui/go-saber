package utils

import (
	"fmt"
	"reflect"
)

// 这里的反射类型指reflect.Type和reflect.Value。
/*
反射的第一定律为：反射可以将接口类型变量转换为反射类型变量。
我们来看reflect.TypeOf的定义，通过该函数，将传入的interface{}类型的变量进
行解析后返回reflect.Type类型
func TypeOf(i interface{}) Type

reflect.ValueOf则是将传入的interface{}类型的变量进行解析后返回reflect.Value类型：
func ValueOf(i interface{}) Value

*/

func Demo8() {
	var a int = 5

	fmt.Printf("type:%T \n", reflect.TypeOf(a))
	fmt.Printf("value:%T \n", reflect.ValueOf(a))
}

/*
反射的第二定律为：反射可以将反射类型变量转换为接口类型。这主要使用了reflect.Value对
象的Interface()方法，可以将一个反射类型变量转换为接口变量
*/

func Demo9() {
	var a int = 5
	valueOfA := reflect.ValueOf(a)

	fmt.Println(valueOfA.Interface())
}

/*
反射的第三定律为：想要使用反射来修改变量的值，其值必须是可写的（CanSet）。这个值必
须满足两个条件：一是变量可以被寻址（CanAddr），二是变量是可导出的（结构体字段名首字母
需大写）。
对于可以被寻址，需要使用reflect.ValueOf()方法对反射对象的地址获取其reflect.Value对象，再
使用Elem()获取实例地址的元素，通过此方式获取的reflect.Value是可以被寻址的
valueOfPerson := reflect.ValueOf(&person)
*/

type OnePerson struct {
	Name string
	age  int `json:"age"`
	string
}

func Demo10() {
	person := OnePerson{"小张", 10, "备注"}
	valueOfPerson := reflect.ValueOf(&person)
	typeOfPerson := reflect.TypeOf(&person)

	for i := 0; i < valueOfPerson.Elem().NumField(); i++ {
		fieldValue := valueOfPerson.Elem().Field(i)
		fieldType := typeOfPerson.Elem().Field(i)
		fmt.Printf("类型名：%v 可以寻址：%v 可以设置：%v \n",
			fieldType.Name, fieldValue.CanAddr(), fieldValue.
				CanSet())
	}

	fmt.Println("修改前：", person)
	// 必须满足可寻址和可导出两个条件才能修改变量值
	valueOfPerson.Elem().Field(0).SetString("xiao zhang")

	fmt.Println("修改后：", person)
}

// 的结构体person中只有Name字段的值可以修改，age字段和匿名字段由于不可导出的原因导致其值不能修改
