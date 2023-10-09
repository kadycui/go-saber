package book

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type OneBook struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func ReadBook() {
	// 读取book.json文件
	data, err := os.ReadFile("tools\\elasticsearch\\book\\book.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析JSON数据
	var books []OneBook
	err = json.Unmarshal(data, &books)
	if err != nil {
		log.Fatal(err)
	}

	// 循环遍历输出书籍信息
	for _, book := range books {
		fmt.Printf("书名：%s\n", book.Title)
		fmt.Printf("作者：%s\n\n", book.Author)
	}
}
