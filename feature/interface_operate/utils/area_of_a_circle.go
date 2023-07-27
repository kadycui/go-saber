package utils

import "fmt"

type area struct {
	r  float64
	pi float64
}

func (a area) circle() float64 {
	var cir = a.pi * a.r * a.r
	return cir
}

func AreaCircle() {

	// 实例化结构体1-var声明
	// var a area
	// a.r = 67.78
	// a.pi = 3.14

	// 实例化结构体2-new关键字
	// a := new(area)
	// a.r = 67.78
	// a.pi = 3.14

	// a := &area{
	// 	67.78, 3.14,
	// }

	// 实例化结构体3-赋值初始化
	// a := area{
	// 	r:  67.78,
	// 	pi: 3.14,
	// }

	// 必须对应结构体定义顺序
	a := area{
		67.78, 3.14,
	}

	var mj = a.circle()
	fmt.Println(mj)

}

/*
var 声明结构体:
	var a area
	为a分配内存, 并0值化


new 关键字声明结构体:
	new(area)
	new返回一个指向area的指针


赋值初始化:
	a := area{67.78, 3.14,}
	为a分配内存, 初始值为67.78, 3.14


	a := &area{67.78, 3.14,}
	new返回一个指向area的指针, 初始值为67.78, 3.14

*/

type Person struct {
	Name string
	Age  int
}

func (p Person) ChangeName(name string) {
	p.Name = name   // 编译器会报错：ineffective assignment to field Person.Name
}

func (p *Person) ChangeAge(age int) {
	p.Age = age
}

func Demo7() {
	person1 := Person{Name: "Tom", Age: 25}
	person2 := person1

	person1.ChangeName("Jerry")
	fmt.Println(person1.Name) // 输出 "Tom"
	fmt.Println(person2.Name) // 输出 "Tom"

	person1.ChangeAge(30)
	fmt.Println(person1.Age) // 输出 30
	fmt.Println(person2.Age) // 输出 25
}
