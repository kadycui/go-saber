package utils

import "fmt"

// var 数组变量名 [数组长度]元素类型

func Demo() {
	var student [3]string // 这是一个空数组
	fmt.Println(student)

	var stu = [3]string{"李白", "杜甫", "苏轼"} // 指定长度数组
	fmt.Println(stu)

	var arr = [...]string{"a", "b", "c", "d"} // 自适应长度数组
	fmt.Println(arr)

	for index, value := range stu {
		fmt.Println("索引k: ", index, " ", "值v: ", value)
	}
}

func Demo1() {
	var num = [...]int{1, 2, 3, 4, 5, 6, 7}
	for k, v := range num {
		fmt.Println("索引k: ", k, " ", "变量值v: ", v)
	}

}
