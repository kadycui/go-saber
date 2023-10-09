package curd

import (
	"context"
	"log"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// SingleUpdate 如果单字段修改可以使用以下修改
func SingleUpdate(client *elastic.Client, index string) {
	_, err := client.Update().
		Index(index).
		Id("2").
		Doc(map[string]interface{}{"name": "Tom"}).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	log.Println("文档修改成功!")

}

type Book struct {
	Id       int     `json:"id"`
	No       int64   `json:"no"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	Rating   float64 `json:"rating"`
	Date     string  `json:"date"`
	Category int     `json:"category"`
	Publish  string  `json:"Publish"`
}

// MultipleUpdate 批量修改, 覆盖
func MultipleUpdate(client *elastic.Client, index string) {
	books := []Book{
		{
			Id:       1,
			No:       1692941134537880000,
			Title:    "忠诚的热爱",
			Author:   "严琦",
			Price:    63.95,
			Rating:   2.3,
			Date:     "2023-05-04T09:10:30+08:00",
			Category: 7,
			Publish:  "河南大学出版社",
		},
		{
			Id:       2,
			No:       1692941134542503200,
			Title:    "诚实的城市",
			Author:   "蒲星",
			Price:    98.71,
			Rating:   6.8,
			Date:     "2023-06-29T02:24:37+08:00",
			Category: 10,
			Publish:  "河南大学出版社",
		},
	}

	// 初始化新的BulkService
	bulkRequest := client.Bulk()

	for _, book := range books {
		// 创建一个更新请求
		doc := elastic.NewBulkUpdateRequest().
			Index(index).
			Type("_doc").
			Id(strconv.Itoa(book.Id)).
			Doc(book)
		// 添加到批量操作
		bulkRequest = bulkRequest.Add(doc)
	}

	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		panic(err)
	}

	log.Println("批量更新完成")

}
