package utils

import "fmt"

/*
类型断言是使用在接口变量上的操作。简单来说，接口类型向普通类型的转换就是类型断言。

t, ok := X.(T)
这里X表示一个接口变量，T表示一个类型（也可为接口类型），这句代码的含义是判断X的
类型是否是T。如果断言成功，则ok为true，t的值为接口变量X的动态值；如果断言失败，则t的值
为类型T的初始值，t的类型始终为T。
*/

func checkType(t interface{}, ok bool) {
	if ok {
		fmt.Println("断言成功")
	} else {
		fmt.Println("断言失败")
	}
	fmt.Printf("变量t的类型=%T, 值=%v \n", t, t)
}

func Demo9() {
	var X interface{} = 1
	fmt.Println("第一次断言:")
	t0, ok := X.(string)
	checkType(t0, ok)

	fmt.Println("第二次断言:")
	t1, ok := X.(float64)
	checkType(t1, ok)

	fmt.Println("第一次断言:")
	t2, ok := X.(int)
	checkType(t2, ok)

}

type Person interface {
	Speak()
}

type NewPerson struct {
	Name string
	Age  int
}

type NewStudent struct {
	Name string
}

func (s NewStudent) Speak() {
	fmt.Println("学生的名字是: ", s.Name)
}

func Demo10() {
	// var stu interface{} = NewStudent{"Tom"}
	// fmt.Println("第一次断言")
	// t0, ok := stu.(string)
	// checkType(t0, ok)
	// fmt.Println("第二次断言")
	// t1, ok := stu.(Person)
	// checkType(t1, ok)

	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "str"
	s[2] = NewPerson{"Tpm", 20}

	for index, data := range s {
		if value, ok := data.(int); ok == true {
			fmt.Printf("s[%d] Type = int, Value = %d\n", index, value)
		}
		if value, ok := data.(string); ok == true {
			fmt.Printf("s[%d] Type = string, Value = %s\n", index, value)

		}
		if value, ok := data.(NewPerson); ok == true {
			fmt.Printf("s[%d] Type = NewPerson, NewPerson = %v, NewPerson.Age = %d\n", index, value.Name, value.Age)
		}
	}

}

// switch-type断言类型

/*
switch value := 接口变量.(type) {
	case类型1:
		// 接口变量是类型1时的处理
	case类型2:
		// 接口变量是类型2时的处理
	case类型3:
		// 接口变量是类型3时的处理
	...
	default:
		// 接口变量不是所有case中所列举的类型时的处理
}
*/

func Demo11() {
	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "str"
	s[2] = NewPerson{"zhangsan", 30}

	for index, data := range s {
		switch value := data.(type) {
		case int:
			fmt.Printf("s[%d] Type = int, Value = %d\n", index, value)
		case string:
			fmt.Printf("s[%d] Type = string, Value = %s\n", index, value)
		case NewPerson:
			fmt.Printf("s[%d] Type = Person, NewPerson.Name = %v, NewPerson.Age = %d\n", index, value.Name, value.Age)
		}
	}
}
