package utils

import "fmt"

/*
接口嵌入，也叫接口组合，在其他语言中，这种接口的组合叫作继承；Go语言舍弃了繁杂
的继承体系，但继承这个特性还是通过接口嵌入得以实现了。接口嵌入也就是指如果一个接口
interface1作为interface2的一个嵌入字段，那么interface2隐式包含了interface1里面所有的方法
*/

type IPerson interface {
	Speak()
}

type IStudent interface {
	IPerson
	Study()
}

type ITeacher interface {
	IPerson
	Teach()
}

type Student struct {
	Name string
}

func (s *Student) Speak() {
	fmt.Println("My name is ", s.Name)
}

func (s *Student) Study() {
	fmt.Println(s.Name, " is studying")
}

type Teacher struct {
	Name string
}

func (t *Teacher) Speak() {
	fmt.Println("My name is", t.Name)
}

func (t *Teacher) Teach() {
	fmt.Println(t.Name, "is teaching")
}

func Demo5() {
	var stu Student = Student{"Tom"}
	var teacher Teacher = Teacher{"Mr. Li"}
	stu.Speak()
	stu.Study()

	teacher.Speak()
	teacher.Teach()
}
