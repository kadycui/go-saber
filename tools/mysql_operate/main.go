package main

import (
	"gostart/mysql_operate/utils"
)

func main() {
	// var sdate string
	// var edate string
	// for i, v := range os.Args {
	// 	if i == 1 {
	// 		sdate = v
	// 	}
	// 	if i == 2 {
	// 		edate = v
	// 	}

	// }
	//var sdate = "2021-12-26 00:00:00"
	//var edate = "2022-12-27 00:00:00"
	// utils.InitDb()
	// utils.CreateTable()
	// utils.InsertData()
	// utils.SelectData()
	// utils.DeleteData()
	// utils.UpdateData()
	// utils.DataRecovery()
	utils.StartPush()
	// utils.RechargeToCsv(sdate, edate)

	//  go build -o test.exe main.go
	// .\test.exe 2023-01-01 2023-01-31

}
