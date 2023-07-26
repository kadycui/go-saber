package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMany() {
	// 创建一个context
	ctx := context.TODO()

	// 设置MongoDB连接选项
	clientOptions := options.Client().ApplyURI("mongodb://admin:123456@localhost:27017/?authSource=admin")

	// 建立连接
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接是否成功
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 断开连接
	defer client.Disconnect(ctx)

	// 选择数据库和集合
	collection := client.Database("mydatabase").Collection("mycollection")

	// 需要插入的多个文档
	persons := []Person{
		{Name: "John Doe", Age: 30, Email: "johndoe@example.com"},
		{Name: "Alice Smith", Age: 25, Email: "alicesmith@example.com"},
		{Name: "Bob Johnson", Age: 35, Email: "bobjohnson@example.com"},
	}

	// 将文档转换为切片
	var documents []interface{}
	for _, person := range persons {
		documents = append(documents, person)
	}

	// 执行插入操作
	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("插入成功")
}
