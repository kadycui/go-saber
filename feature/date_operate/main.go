package main

import (
	"fmt"
	"time"
)

//var WeekDayMap = map[string]string{
//	"Monday":    "1",
//	"Tuesday":   "2",
//	"Wednesday": "3",
//	"Thursday":  "4",
//	"Friday":    "5",
//	"Saturday":  "6",
//	"Sunday":    "7",
//}

func main() {
	now := time.Now().Unix()
	fmt.Println("now是时间戳", now)
	StartTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local).Unix()
	EndTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 59, time.Local).Unix()

	StartTimeStr := time.Unix(StartTime, 0).Format("2006-01-02 15:04:05")
	EndTimeStr := time.Unix(EndTime, 0).Format("2006-01-02 15:04:05")
	fmt.Println("StartTime是时间戳", StartTime, "| StartTimeStr是时间字符串", StartTimeStr)
	fmt.Println("EndTime是时间戳", EndTime, "| EndTimeStr是时间字符串", EndTimeStr)

	day := "2022-11-01 00:00:00"
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", day, time.Local)
	fmt.Println("t1是时间类型", t1)

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", StartTimeStr, time.Local)
	fmt.Println("t2是时间类型", t2)

	//week := WeekDayMap[t1.Weekday().String()]
	//fmt.Println("当前时间是周", week)

	date := t1.Format("2006-01-02")
	fmt.Println("当前的日期是", date)
	fmt.Println("================================================")
	for {
		fmt.Println("startT1=", t1.Format("2006-01-02 15:04:05"))
		//fmt.Println("今天周", WeekDayMap[t1.Weekday().String()])
		fmt.Println("日期是", t1.Format("2006-01-02"))
		formatT1 := time.Date(t1.Year(), t1.Month(), t1.Day(), 23, 59, 59, 59, time.Local).Unix()
		fmt.Println("endT1=", time.Unix(formatT1, 0).Format("2006-01-02 15:04:05"))
		fmt.Println("================================================")

		t1 = t1.AddDate(0, 0, 1)

		if t1.After(t2) || t1.Equal(t2) {
			break

		}

	}

}
