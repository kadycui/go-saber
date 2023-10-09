package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间的时间戳（秒）
	timestamp := time.Now().Unix()

	fmt.Println(timestamp)



	// 获取当前时间的纳秒级时间戳
	timestampNano := time.Now().UnixNano()
	fmt.Println(timestampNano)


	// 将纳秒级时间戳转换为以毫秒为单位的时间戳
	timestampMillis := timestampNano / int64(time.Millisecond)

	fmt.Println(timestampMillis)

}
