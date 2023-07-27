package main

import "gostart/csv_operate/utils"

func main() {

	var data = [][]string{{"tom", "18", "beijing"}, {"jon", "19", "shanghai"}}
	var filePath1 = "D:\\GoCode\\gostart\\csv_operate\\utils\\read_data.csv"
	var filePath2 = "D:\\GoCode\\gostart\\csv_operate\\utils\\write_data.csv"

	// utils.ReadDataLine(filePath)
	utils.ReadDataAll(filePath1)

	utils.WriteData(filePath2, data)

}
