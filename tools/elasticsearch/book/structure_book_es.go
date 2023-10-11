package book

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

type Book struct {
	Id       int     `json:"id"`
	No       int64   `json:"no"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	Rating   float64 `json:"rating"`
	Date     string  `json:"date"`
	Category int     `json:"category"`
	Publish  string  `json:"publish"`
}

func MakeBookData() {
	// 创建Elasticsearch客户端
	client, err := elastic.NewClient(
		elastic.SetURL("http://10.16.168.61:9200/"),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个book索引（如果还不存在）
	indexName := "book"
	// 检查索引是否存在
	exists, err := client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Fatalf("检查索引是否存在失败：%s", err)
	}

	if exists {
		// 索引存在，跳过检查
		fmt.Printf("索引%s存在\n", indexName)
	} else {
		// 索引不存在，创建索引
		createIndex, err := client.CreateIndex(indexName).BodyString(`
			{
				"settings": {
					"number_of_shards": 1,
					"number_of_replicas": 0
				}
			}
		`).Do(context.Background())
		if err != nil {
			log.Fatalf("创建索引失败：%s", err)
		}
		if !createIndex.Acknowledged {
			log.Fatalf("创建索引未被确认")
		}
		fmt.Printf("索引%s已创建\n", indexName)
	}

	// 读取book.json文件
	data, err := os.ReadFile("tools\\elasticsearch\\book\\book.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析JSON数据
	var onebooks []OneBook
	err = json.Unmarshal(data, &onebooks)
	if err != nil {
		log.Fatal(err)
	}
	num := 20000
	// 循环遍历输出书籍信息
	for i := 1; i <= num; i++ {
		// fmt.Printf("书名：%s\n", onebook.Title)
		// fmt.Printf("作者：%s\n\n", onebook.Author)
		// 索引一本新书
		book := Book{
			Id:       i,
			No:       time.Now().UnixNano(),
			Title:    BookName(),
			Author:   GetFullName(),
			Price:    Price(),
			Rating:   Rating(),
			Date:     RandomTimeStr(),
			Category: Category(),
			Publish:  Publish(),
		}

		indexResp, err := client.Index().
			Index(indexName).
			Id(strconv.Itoa(i)).
			BodyJson(book).
			Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Indexed book: %s\n", indexResp.Id)

	}

	// 获取一本书
	// getResp, err := client.Get().
	// 	Index(indexName).
	// 	Id("1").
	// 	Do(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if getResp.Found {
	// 	var book Book
	// 	err := json.Unmarshal(getResp.Source, &book)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Got book: %+v\n", book)
	// }

	// 更新一本书
	// book.Title = "The Go Programming Language (Updated)"
	// updateResp, err := client.Update().
	// 	Index(indexName).
	// 	Id("1").
	// 	Doc(book).
	// 	Do(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Updated book: %s\n", updateResp.Id)

	// 删除一本书
	// deleteResp, err := client.Delete().
	// 	Index(indexName).
	// 	Id("1").
	// 	Do(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Deleted book: %s\n", deleteResp.Id)
}
