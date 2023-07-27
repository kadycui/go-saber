package utils

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CRUD() {
	// 创建一个context
	ctx := context.TODO()

	// 设置MongoDB连接选项，包括用户名、密码和认证数据库
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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
	collection := client.Database("test").Collection("persons")

	// 创建文档
	person := Person{
		Name:  "kadycui",
		Age:   28,
		Email: "kadycui@qq.com",
	}

	// 插入文档
	result, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}

	// 获取插入的文档ID
	insertedID := result.InsertedID
	fmt.Println("Inserted ID:", insertedID)

	// 根据条件查找单个文档
	var resultPerson Person
	err = collection.FindOne(ctx, bson.M{"name": "kadycui"}).Decode(&resultPerson)
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Println("Name:", resultPerson.Name)
	fmt.Println("Age:", resultPerson.Age)
	fmt.Println("Email:", resultPerson.Email)

	// 更新文档
	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "age", Value: 35}}},
	}
	_, err = collection.UpdateOne(ctx, bson.M{"name": "John Doe"}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document updated")

	// 删除文档
	_, err = collection.DeleteOne(ctx, bson.M{"name": "John Doe"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document deleted")
}
