package utils

import "fmt"

// 结构体内嵌

type Student struct {
	sid  int
	name string
	age  int
}

type OneStudent struct {
	Student
	admission_time string
}

type TwoStudent struct {
	Student
	addr string
}

func Demo8() {
	stu1 := &OneStudent{}
	stu2 := &TwoStudent{}
	fmt.Println(stu1)
	fmt.Println(stu2)

}

type Movie struct {
	title string
	actor string
	id    int
	num   int
}

type MovieBorrow struct {
	Movie
	borrowTime string
}

type MovieNotBorrow struct {
	Movie
	readTime string
}

func Demo9() {

	// 初始化结构体内嵌

	// 1-键值对
	movieBorrow := &MovieBorrow{
		Movie: Movie{
			"西游记",
			"六小龄童",
			10001,
			30,
		},
		borrowTime: "50",
	}
	fmt.Println(movieBorrow)

	// 2-.属性
	movieNotBorrow := &MovieNotBorrow{}
	movieNotBorrow.title = "水浒传"
	movieNotBorrow.actor = "李雪健"
	movieNotBorrow.id = 10002
	movieNotBorrow.num = 60
	movieNotBorrow.readTime = "70"

	fmt.Println(movieNotBorrow)
}

// 内嵌匿名结构体

type BookBorrow2 struct {
	Book2 struct { // 匿名结构体
		title  string
		author string
		num    int
		id     int
	}
	borrowTime string
}

func Demo10() {
	// 匿名结构体实例化
	bookBorrow2 := &BookBorrow2{
		Book2: struct { // 声明类型
			title  string
			author string
			num    int
			id     int
		}{title: "go language", author: "Tom", num: 20, id: 1003},
		borrowTime: "60",
	}
	fmt.Println(bookBorrow2)
}
