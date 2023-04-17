package utils

import (
	"fmt"
	"reflect"
)

type RePerson struct {
	Name string
	Age  int `json:"age"`
	string
}

func Demo5() {
	person := RePerson{"小张", 10, "备注"}
	typeOfPerson := reflect.TypeOf(person)

	// 遍历所有结构体成员获取字段信息
	fmt.Println("遍历结构体")
	for i := 0; i < typeOfPerson.NumField(); i++ {
		field := typeOfPerson.Field(i)
		fmt.Printf("字段名：%v 字段标签：%v 是否是匿名字段：%v \n", field.Name, field.Tag, field.Anonymous)
	}

	// 通过字段名获取字段信息
	if field, ok := typeOfPerson.FieldByName("Age"); ok {
		fmt.Println("通过字段名")
		fmt.Printf("字段名：%v, 字段标签中json：%v \n", field.Name,
			field.Tag.Get("json"))
	}

	// 通过下标获取字段信息
	field := typeOfPerson.FieldByIndex([]int{1})
	fmt.Println("通过下标")
	fmt.Printf("字段名：%v, 字段标签：%v \n", field.Name, field.Tag)
}
