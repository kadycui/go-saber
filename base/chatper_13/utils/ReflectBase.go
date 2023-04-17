package utils

/*
Go语言提供了一种机制：在运行时更新变量和检查它们的值、调用它们的方法和它们支持的
内在操作，但是在编译时并不知道这些变量的具体类型。这种机制被称为反射。
*/

import (
	"fmt"
	"reflect"
)

func Demo() {
	// 使用reflect.TypeOf()获取了变量a的值类型，使用reflect.ValueOf()获取了变量a的原始值
	var a interface{} = "这是一个字符串!"
	typeOfa := reflect.TypeOf(a)
	fmt.Println("变量a的类型为:" + typeOfa.Name())
	vfa := reflect.ValueOf(a)
	if typeOfa.Kind() == reflect.String {
		fmt.Println("变量a的值为：" + vfa.String())

	}
}

func Demo1() {
	var a int
	typeofa := reflect.TypeOf(a)
	if typeofa.Kind() == reflect.Int {
		fmt.Printf("变量a的Kind是int")

	} else {
		fmt.Printf("变量a的Kind不是int")
	}
}

/*
Go语言中使用reflect.TypeOf()来获取变量对象的类型信息，返回值的反射类型对象（reflect.Type)
var typeOfNum reflect.Type = reflect.TypeOf(num)

获取类型名称可以使用reflect.Type中的Name()方法；获取类型的种类Kind，可以使用对应的Kind()方法；
对于指针类型变量，可以使用reflect.Type中的Elem()来反射获取指针指向的元素类型
*/

//定义了int种类的Number类型和一个Person结构体类型。checkType()函数对传
//入的reflect.Type变量进行判断，如果变量的种类为指针类型，就输出类型种类，并获取其指向的元
//素，输出类型名和类型Kind。

type Number int

type Person struct {
}

func checkType(t reflect.Type) {
	if t.Kind() == reflect.Ptr {
		fmt.Printf("变量的类型名称%v, 指向的变量为:", t.Kind())
		t = t.Elem()
	}

	fmt.Printf("变量的类型名称 => %v, 类型种类 => %v \n", t.Name(), t.Kind())

}

func Demo2() {
	var num Number = 10
	typeOfNum := reflect.TypeOf(num)
	fmt.Println("typeOfNum")
	checkType(typeOfNum)

	var person Person
	typeOfPerson := reflect.TypeOf(person)
	fmt.Println("typeOfPerson")
	checkType(typeOfPerson)

	typeOfPersonPtr := reflect.TypeOf(&person)
	fmt.Println("typeOfPersonPtr")
	checkType(typeOfPersonPtr)

}
