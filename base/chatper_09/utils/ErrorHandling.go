package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"
)

func Demo1() {
	f, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err) // open test.txt: The system cannot find the file specified.

	} else {
		fmt.Println(f)
	}
}

// Go语言的error类型实际上是一个接口，定义如下：
/*
type error interface {
 Error() string
}
*/

func Demo2() {
	err := errors.New("This is a error")
	var err2 error
	fmt.Println(err)  // This is a error
	fmt.Println(err2) // <nil>
}

func Demo3() {
	if _, _, line, ok := runtime.Caller(0); ok == true {
		err := fmt.Errorf("****Line %d error****", line)
		fmt.Println(err.Error())
	}
}
