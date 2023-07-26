package utils

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func MakeMongoData() {
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
	collection := client.Database("test").Collection("persons")

	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 插入10000条随机数据
	for i := 0; i < 10000; i++ {
		// 随机生成姓名和邮箱
		name := uuid.New().String()[:8]
		email := fmt.Sprintf("%s@example.com", uuid.New().String()[:8])

		// 随机生成年龄
		age := rand.Intn(100)

		// 创建Person对象
		person := Person{
			Name:  name,
			Age:   age,
			Email: email,
		}

		// 插入文档
		_, err := collection.InsertOne(ctx, person)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("插入成功")
}
