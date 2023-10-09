package main

import (
	"fmt"

	"github.com/kadycui/go-saber/tools/json_type_conversion/utils"
)

func main() {

	utils.StructToJson()
	utils.MapToJson()
	utils.IntToJson()
	utils.SliceToJson()
	utils.JsonToStruct()
	utils.JsonToMap()

	fmt.Println("-------------------------------------------")

	utils.Simplejson()

}
