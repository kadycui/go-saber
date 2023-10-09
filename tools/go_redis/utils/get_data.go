package utils

import "fmt"

func GetData() {
	client := NewRedisClient()
	key := "449839453"
	val, err := client.Get(key)
	if err != nil {
		fmt.Printf("链接redis错误, 错误信息: %v", err)
	} else {
		fmt.Println(val)
	}

}
