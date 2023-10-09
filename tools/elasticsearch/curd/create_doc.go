package curd

import (
	"context"
	"log"

	"github.com/olivere/elastic/v7"
)

type User struct {
	name string
	age  int
}

// StructForm 结构体形式创建索引
func StructForm(client *elastic.Client, index string) {
	u := User{
		name: "张三",
		age:  18,
	}

	_, err := client.Index().
		Index(index).
		Id("1").
		BodyJson(u).Do(context.Background())

	if err != nil {
		panic(err)
	}
	log.Printf("索引 %s中的文档创建成功", index)

}

// StrForm 字符串形式创建
func StrForm(client *elastic.Client, index string) {
	u := `{"name":"", "age":21}`

	_, err := client.Index().
		Index(index).
		Id("3").
		BodyJson(u).Do(context.Background())

	if err != nil {
		panic(err)
	}
	log.Printf("索引 %s中的文档创建成功", index)

}
