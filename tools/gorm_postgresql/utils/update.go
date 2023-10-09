package utils

import (
	"log"

	"github.com/kadycui/go-saber/tools/gorm_postgresql/model"
)

func UpdateData() {
	db := GetDB()
	result := db.Model(&model.Book{}).Where("author = ?", "满嘉").Updates(model.Book{Publish: "东莞出版社"})
	if result.Error != nil {
		log.Println("查询错误")
	}

}
