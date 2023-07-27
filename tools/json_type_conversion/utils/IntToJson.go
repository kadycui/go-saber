package utils

import (
	"encoding/json"
	"fmt"
)

func IntToJson() {
	var age = 26

	data, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json.Marshal failed, err:", err)
		return

	}

	fmt.Printf("%s\n", string(data))
}
