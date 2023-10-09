package utils

import (
	"fmt"
)

func Ping() {
	// rdb := GetDb()
	rdb := NewRedisClient()
	pong, err := rdb.Ping()
	if err != nil {
		fmt.Printf("链接redis错误, 错误信息: %v", err)
	} else {
		fmt.Println(pong)
		fmt.Println("成功链接!")

	}

}
