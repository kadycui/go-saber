package utils

import (
	"log"

	"github.com/kadycui/go-saber/tools/gorm_postgresql/model"
)

func DeleteData() {
	var books []model.Book
	db := GetDB()
	result := db.Where("category > ?", 8).Delete(&books)
	if result.Error != nil {
		log.Println("查询错误")
	}

}
