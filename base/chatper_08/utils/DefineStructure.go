package utils

import "fmt"

/*
type 结构体名 struct {
 成员变量1 类型1
 成员变量2 类型2
 成员变量3 类型3
 ...
}


◇ 结构体名：同一个包内结构体名不能重复。
◇ 成员名：同一个结构体内，成员名不能重复。
◇ 类型1、类型2……：表示结构体成员变量的类型。
◇ 同类型的成员名可以写在同一行。
◇ 当结构体、方法名或变量名的首字母为大写时（可被导出），就可以在当前包外进行访问。
*/

type Book struct {
	title  string
	author string
	num    int
	id     int
}

type NewBook struct {
	title, author string
	num, id       int
}

func Demo1() {
	var book1 Book
	fmt.Println(book1)
}

func Demo2() {
	book1 := new(Book)
	fmt.Println(book1)
}

func Demo3() {
	book1 := &Book{}
	book1.title = "Go 语言"
	book1.author = "Tom"
	book1.num = 20
	book1.id = 152368
	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)
}

/*
结构体实例 := 结构体类型{
 成员变量1:值1,
 成员变量2:值2,
 成员变量3:值3,
}
*/

func Demo4() {
	book1 := &Book{
		title:  "Go 语言",
		author: "Tom",
		num:    200,
		id:     12345,
	}

	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)
}

/*
结构体实例 := 结构体类型{
 值1,
 值2,
 值3,
}
◇ 使用这种方式初始化结构体必须初始化所有的成员变量。
◇ 值的填充顺序必须和结构体成员变量声明顺序保持一致。
◇ 该方式与键值对的初始化方式不能混用。

*/

func Demo5() {
	book1 := &Book{
		"GO语言",
		"Tom",
		20,
		12345466,
	}

	fmt.Println("title:", book1.title)
	fmt.Println("author:", book1.author)
	fmt.Println("num:", book1.num)
	fmt.Println("id:", book1.id)
}
