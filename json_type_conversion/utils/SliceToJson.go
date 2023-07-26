package utils

import (
	"encoding/json"
	"fmt"
)

func stj(s []map[string]interface{}) string {

	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(data)

}

func SliceToJson() {

	var s []map[string]interface{}

	tempMap := make(map[string]interface{})

	tempMap["username"] = "kadycui"
	tempMap["age"] = 25
	tempMap["sex"] = "男"

	s = append(s, tempMap)

	tempMap = make(map[string]interface{})
	tempMap["username"] = "LilIli"
	tempMap["age"] = 24
	tempMap["sex"] = "女"

	s = append(s, tempMap)

	str := stj(s)

	fmt.Println(str)

}
