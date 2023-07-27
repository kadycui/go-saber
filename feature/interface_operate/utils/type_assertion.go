package utils

import "fmt"

func judgeType1(q interface{}) {
	temp, ok := q.(string)
	if ok {
		fmt.Println("类型转换成功", temp)
	} else {
		fmt.Println("类型转换失败", temp)
	}

}

func judgeType2(q interface{}) {
	switch i := q.(type) {
	case string:
		fmt.Println("这是一个字符串", i)
	case int:
		fmt.Println("这是一个数字", i)
	case bool:
		fmt.Println("这是一个布尔类型", i)
	default:
		fmt.Println("这是未知类型")
	}
}

func Demo5() {
	a := "12"
	judgeType1(a)

	b := make([]int, 5, 20)
	judgeType2(b)
}
