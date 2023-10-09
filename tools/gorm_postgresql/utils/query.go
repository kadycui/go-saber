package utils

import (
	"fmt"
	"log"

	"github.com/kadycui/go-saber/tools/gorm_postgresql/model"
)

func FindAll() {

	var books []model.Book
	db := GetDB()

	result := db.Find(&books)
	if result.Error != nil {
		log.Println("查询错误")
	}
	fmt.Println("符合条件的数据:", len(books))
	for _, book := range books {
		fmt.Println(book.Id, book.Title, book.Author)
	}
}

func WhereQuery() {
	var books []model.Book
	db := GetDB()
	result := db.Where("category > ?", 8).Find(&books)
	if result.Error != nil {
		log.Println("查询错误")
	}
	fmt.Println("符合条件的数据:", len(books))
	for _, book := range books {
		fmt.Println(book.Id, book.Title, book.Category)
	}

}
