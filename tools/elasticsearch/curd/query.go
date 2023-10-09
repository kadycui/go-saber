package curd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

func IdQuery(client *elastic.Client, index string, id string) {
	ret, err := client.Get().Index("book").
		Id(id).Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("id:%s \n Source:%s ", ret.Id, string(ret.Source))

}

// MatchQuery 匹配查询Match Query
func MatchQuery(client *elastic.Client) {
	ctx := context.Background()
	// 构建查询条件
	query := elastic.NewMatchQuery("title.keyword", "人民出版社")

	// 执行查询
	result, err := client.Search().
		Index("book").
		Query(query).
		Do(ctx)
	if err != nil {
		// 处理错误
		fmt.Println("查询数据为空!")
		return
	}

	// 处理查询结果
	for _, hit := range result.Hits.Hits {
		fmt.Printf("Document ID: %s\n", hit.Id)
		docJSON, err := json.Marshal(hit)
		if err != nil {
			fmt.Println("格式转化异常!")
			return
		}
		// 输出 JSON 格式化后的结果
		fmt.Println(string(docJSON))
	}

}

// MatchQueryToJson 匹配查询Match Query转Json
func MatchQueryToJson(client *elastic.Client) {
	ctx := context.Background()
	// 构建查询条件
	query := elastic.NewMatchQuery("publish.keyword", "人民出版社")

	// 执行查询
	result, err := client.Search().
		Index("book").
		Query(query).
		Do(ctx)
	if err != nil {
		// 处理错误
		fmt.Println("查询数据为空!")
		return
	}

	// 解组每个文档结果为 Book 结构体
	var books []Book
	for _, hit := range result.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			// 处理错误
			return
		}
		fmt.Println(book)
		books = append(books, book)
	}

	// 输出解组后的结果
	for _, book := range books {
		fmt.Printf("Book ID: %d\n", book.Id)
		fmt.Printf("Book Title: %s\n", book.Title)
		// 其他输出...
	}

}

// RangeQuery 范围查询
func RangeQuery(client *elastic.Client) {
	ctx := context.Background()

	// 构建查询寻条件
	query := elastic.NewRangeQuery("category").Gte(8).Lt(10)

	// 执行查询
	result, err := client.Search().
		Index("book").
		Query(query).
		Do(ctx)
	if err != nil {
		// 处理错误
		fmt.Println("查询数据为空!")
		return
	}

	// 解组每个文档结果为 Book 结构体
	var books []Book
	for _, hit := range result.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			// 处理错误
			return
		}
		fmt.Println(book)
		books = append(books, book)
	}

	// 输出解组后的结果
	for _, book := range books {
		fmt.Printf("Book ID: %d\n", book.Id)
		fmt.Printf("Book Category: %d\n", book.Category)
		// 其他输出...
	}

}

// BoolQuery 布尔查询
func BoolQuery(client *elastic.Client) {
	ctx := context.Background()
	query := elastic.NewBoolQuery().
		Must(elastic.NewTermQuery("category", 5)).
		MustNot(elastic.NewTermQuery("author", "申永"))

	// 执行查询
	result, err := client.Search().
		Index("book").
		Query(query).
		Do(ctx)
	if err != nil {
		// 处理错误
		fmt.Println("查询数据为空!")
		return
	}

	// 解组每个文档结果为 Book 结构体
	var books []Book
	for _, hit := range result.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			// 处理错误
			return
		}
		fmt.Println(book)
		books = append(books, book)
	}

	// 输出解组后的结果
	for _, book := range books {
		fmt.Printf("Book ID: %d\n", book.Id)
		fmt.Printf("Book Category: %d\n", book.Category)
		// 其他输出...
	}
}

// PrefixQuery 前缀查询
func PrefixQuery(client *elastic.Client) {
	ctx := context.Background()
	// query := elastic.NewPrefixQuery("title", "忧")

	query := elastic.NewWildcardQuery("title", "忧*")

	// 执行查询
	result, err := client.Search().
		Index("book").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		// 处理错误
		fmt.Println("查询数据为空!")
		return
	}

	// 解组每个文档结果为 Book 结构体
	var books []Book
	for _, hit := range result.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			// 处理错误
			return
		}
		fmt.Println(book)
		books = append(books, book)
	}

	// 输出解组后的结果
	for _, book := range books {
		fmt.Printf("Book ID: %d\n", book.Id)
		fmt.Printf("Book Category: %d\n", book.Category)
		// 其他输出...
	}
}

// PrefixQuery2 前缀查询
func PrefixQuery2(client *elastic.Client) {
	ctx := context.Background()
	// query := elastic.NewPrefixQuery("title", "忧")

	query := elastic.NewMatchPhrasePrefixQuery("title", "忧")

	// 执行查询
	searchResult, err := client.Search().
		Index("book").
		Query(query).
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Fatalf("查询失败：%s", err)
	}

	// 处理查询结果
	if searchResult.Hits.TotalHits.Value > 0 {
		fmt.Printf("查询到 %d 条结果：\n", searchResult.Hits.TotalHits.Value)
		for _, hit := range searchResult.Hits.Hits {
			fmt.Printf("文档 ID: %s，文档内容: %s\n", hit.Id, hit.Source)
		}
	} else {
		fmt.Println("未找到匹配的结果")
	}
}
