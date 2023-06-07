package main

import (
	"gostart/gorm_operate/utils"
)

func main() {
	db := utils.GetDB()

	// 批量插入
	users := []utils.User{
		{Name: "Tom", Age: 30, Email: "tom@example.com"},
		{Name: "Jerry", Age: 28, Email: "jerry@example.com"},
		{Name: "Mike", Age: 35, Email: "mike@example.com"},
	}

	result := db.Create(&users)
	if result.Error != nil {
		panic(result.Error)
	}

}
