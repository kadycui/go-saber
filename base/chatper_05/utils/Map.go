package utils

import (
	"fmt"
)

// var map [键类型]值类型

func Demo31() {
	//var studentScoreMap map[string]int
	//fmt.Println(studentScoreMap)

	var studentScoreMap = map[string]int{
		"Tom": 80,
		"Jay": 90,
		"May": 93,
	}
	fmt.Println(studentScoreMap)
}

func Demo32() {
	/*
		make(map[键类型]值类型,map容量)

		使用make()函数初始化map时可以不指定map容量，但是对于map的多次扩充会造成性能损耗。
		cap()函数只能用于获取切片的容量，无法获得map的容量，因此可以通过len()函数获取map的当前长度
	*/
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)

	studentScoreMap["Tom"] = 80
	studentScoreMap["Jay"] = 90
	studentScoreMap["May"] = 92
	fmt.Println("map长度为：", len(studentScoreMap))
	fmt.Println(studentScoreMap)

	// 遍历映射
	for key, value := range studentScoreMap {
		fmt.Println("映射的键是:", key, "", "映射的值是: ", value)
	}

	//只遍历键
	for k := range studentScoreMap {
		fmt.Println(k)
	}

	// 只遍历值
	for _, v := range studentScoreMap {
		fmt.Println(v)
	}
}

func Demo33() {
	// 映射中删除键值对 delete(map,键)
	var studentScoreMap map[string]int
	studentScoreMap = make(map[string]int)
	studentScoreMap["Tom"] = 80
	studentScoreMap["Jay"] = 90
	studentScoreMap["May"] = 92
	fmt.Println("map长度为：", len(studentScoreMap))
	fmt.Println(studentScoreMap)

	delete(studentScoreMap, "Tom")
	fmt.Println("map长度为：", len(studentScoreMap))
	fmt.Println(studentScoreMap)

}

func readMap(GoMap map[int]int, key int) int {
	fmt.Println("我是readMap")
	return GoMap[key]

}

func writeMap(GoMap map[int]int, key int, value int) {
	fmt.Println("我是writeMap")
	GoMap[key] = value
}

func Demo34() {
	fmt.Println("333333333333")
	GoMap := make(map[int]int)
	for i := 0; i < 10000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)
	}
	fmt.Println("Done")
}
