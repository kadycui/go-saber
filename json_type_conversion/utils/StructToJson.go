package utils

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// 结构体转json
func StructToJson() {
	user := User{
		UserName: "kady",
		NickName: "kk",
		Age:      28,
		Birthday: "1994-01-01",
		Sex:      "男",
		Email:    "kadycui@qq.com",
		Phone:    "18236322325",
	}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal failed, err:", err)
		return

	}

	fmt.Printf("%s\n", string(data))
}
