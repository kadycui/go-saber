package utils

import (
	"fmt"
	"regexp"
)

func Demo3() {
	// FindStringIndex函数接收一个目标字符串，并返回第一个匹配的起始位置和结束位置

	testString := "hello golang"

	// MustCompile函数：若正则表达式未通过编译，则引发panic
	re := regexp.MustCompile(`(\w)+`)

	res := re.FindStringIndex(testString)
	fmt.Println(res) // [0 5]

}

func Demo4() {
	// ReplaceAllString函数返回第一个参数的拷贝，将第一个参数中所有re的匹配结果都替换为repl
	testString := "hello golang"

	re := regexp.MustCompile(`o`)
	res := re.ReplaceAllString(testString, "O")
	fmt.Println(res) // hellO gOlang

}
