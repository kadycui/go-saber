package utils

import "fmt"

/*
◇ 地址：切片的地址一般指切片中第一个元素所指向的内存地址，用十六进制表示。
◇ 长度：切片中实际存在元素的个数。
◇ 容量：从切片的起始元素开始到其底层数组中的最后一个元素的个数。
*/

func Demo2() {
	// 从数组生成一个新切片
	var stu = [...]string{"Tom", "Jack", "Mary"}
	var ss = stu[1:2]

	fmt.Println("stu数组:", stu)
	fmt.Println("stu数组元素[1]的地址为:", &stu[1]) //取stu[1]元素的地址
	fmt.Println("stu切片:", ss)

	fmt.Println("stu切片[0]的地址为:", &ss[0])
	fmt.Println("stu切片的长度为:", len(ss))
	fmt.Println("stu切片的容量为:", cap(ss))

}

func Demo3() {
	// 从切片生成切片
	var student = [...]string{"Tom", "Ben", "Peter"}
	var student1 = student[1:3]
	var student2 = student1[0:1]
	fmt.Println("student数组：", student[:])
	fmt.Println("student1切片：", student1[:])
	fmt.Println("student2切片：", student2[:])
	fmt.Println("student数组地址为", &student[1])
	fmt.Println("student1切片地址为", &student1[0])
	fmt.Println("student2切片地址为", &student2[0])
	fmt.Println("student1切片长度为：", len(student1))
	fmt.Println("student1切片容量为：", cap(student1))
	fmt.Println("student2切片长度为：", len(student2))
	fmt.Println("student2切片容量为：", cap(student2))
}

func Demo4() {
	var student []int // 声明一个空切片
	fmt.Println("student切片：", student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))
	fmt.Println("判定student切片是否为空：", student == nil)

	var student1 = []string{"Tom", "Ben", "Peter"} // 初始化一个切片
	fmt.Println("student1切片：", student1)
	fmt.Println("student1切片长度：", len(student1))
	fmt.Println("student1切片容量：", cap(student1))
	fmt.Println("判定student1切片是否为空：", student1 == nil)

}

func Demo5() {
	// make([]元素类型,切片长度,切片容量)
	var student []int
	student = make([]int, 2, 10)
	fmt.Println("student切片：", student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))
	fmt.Println("判定student切片是否为空：", student == nil)

}

func Demo6() {
	stu := make([]int, 1, 1)
	for i := 0; i < 8; i++ {
		stu = append(stu, i)
		fmt.Println("当前切片长度：", len(stu), "当前切片容量：", cap(stu))
	}

	var stu2 = [...]string{"Tom", "Ben", "Peter"}
	var stu3 = stu2[0:1] //从stu2数组生成切片student1
	fmt.Println("stu2数组：", stu2)
	fmt.Println("stu3切片：", stu3)
	stu3 = append(stu3, "Danny") //对stu3切片的元素添加，会 覆盖引用数组对应的元素
	fmt.Println("扩充Danny后的stu3切片：", stu3, ",切片长度为：", len(stu3), ",切片容量为：", cap(stu3))
	fmt.Println("扩充Danny后的stu2数组：", stu2)

}

func Demo7() {
	// 切片删除
	var student = []string{"Tom", "Ben", "Peter", "Danny"}

	//student = append(student[0:1], student[2:]...)
	student = append(student[0:1], student[2], student[3])

	fmt.Println("student切片：", student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))

	// 清空切片
	student = student[0:0]
	fmt.Println("student切片：", student)
	fmt.Println("student切片长度：", len(student))
	fmt.Println("student切片容量：", cap(student))

}

func Demo8() {
	// 切片遍历
	var student = []string{"Tom", "Ben", "Peter", "Danny"}
	for index, value := range student {
		fmt.Println("切片的索引是:", index, "", "切片的值是:", value)
	}
}
