package utils

import (
	"fmt"
	"regexp"
)

func Demo1() {
	// MarchString函数接收一个要查找的正则表达式和目标字符串，并根据匹配结果返回true或 false
	testString := "hello world"
	matchString := "Hello"
	match, err := regexp.MatchString(matchString, testString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(match)

}

func Demo2() {
	// 如果要以不区分大小写的方式查找，必须修改正则表达式，使用一种特殊的语法
	// matchString := "(?i)hello"
	testsString := "Hello Golang"
	matchString := "(?i)hello"

	match, err := regexp.MatchString(matchString, testsString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(match)

}
