package utils

import (
	"fmt"
	"reflect"
)

/*
获取结构体成员的值，我们仍需要使用reflect.ValueOf()方法获取到reflect.Value对象，使用对象
的NumField()方法可以获取结构体成员的数量，使用Field()则可以根据索引返回对应结构体字段的
reflect.Value反射值类型。
*/

type StrPerson struct {
	Name string
	Age  int `json:"age"`
	string
}

// Field()、FieldByName()、FieldByIndex()三个方法来获取结构体成员的值类型对象

func Demo6() {
	person := StrPerson{"小张", 10, "备注"}
	valueOfPerson := reflect.ValueOf(person)

	fmt.Println("person的字段数量:%v\n", valueOfPerson.NumField())

	// 通过下标访问获取字段值
	fmt.Println("Field")
	field := valueOfPerson.Field(1)
	fmt.Println("字段值: %v \n", field.Int())

	// 通过字段名获取字段值
	field = valueOfPerson.FieldByName("Age")
	fmt.Println("FieldByName")
	fmt.Printf("字段值：%v \n", field.Interface())

	// 通过下标索引获取字段值
	field = valueOfPerson.FieldByIndex([]int{0})
	fmt.Println("FieldByIndex")
	fmt.Printf("字段值：%v \n", field.Interface())
}

// 结构体方法

func (p StrPerson) GetName() {
	fmt.Println(p.Name)
}

func Demo7() {
	person := StrPerson{"小李", 23, "测试"}
	valueOfPerson := reflect.ValueOf(person)

	// 根据名字获取方法
	f := valueOfPerson.MethodByName("GetName")

	// 执行结构体方法
	f.Call([]reflect.Value{})
}
