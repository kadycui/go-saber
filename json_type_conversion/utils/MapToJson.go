package utils

import (
	"encoding/json"
	"fmt"
)

func mtj(temMap *map[string]interface{}) string {

	data, err := json.Marshal(temMap)
	if err != nil {
		panic(err)
	}

	return string(data)

}

func MapToJson() {
	tempMap := make(map[string]interface{})

	tempMap["username"] = "kadycui"
	tempMap["age"] = 25
	tempMap["sex"] = "男"

	str := mtj(&tempMap)

	fmt.Println(str)

}
