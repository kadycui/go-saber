package utils

import (
	"encoding/json"
	"fmt"
)

// map转json
func mapToJSON(tempMap *map[string]interface{}) string {

	data, err := json.Marshal(tempMap)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func jtm(str string) {
	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	fmt.Println(tempMap)

}

func JsonToMap() {

	tempMap := make(map[string]interface{})
	tempMap["username"] = "itbsl"
	tempMap["age"] = 18
	tempMap["sex"] = "male"

	str := mapToJSON(&tempMap)

	jtm(str)

}
