package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPath() string {
	// 获取当前执行文件绝对路径（go run）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}
	dir := path.Dir(filename)
	return dir
}

func Server2Mongo() {
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
	collection := client.Database("gms").Collection("servers")

	fmt.Println(collection)

	// 读取JSON数据
	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "../conf/servers.json")
	file, err := os.ReadFile(p2)
	if err != nil {
		log.Fatal(err)
	}

	var servers []Server
	err = json.Unmarshal([]byte(file), &servers)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		// fmt.Println(server.ID, server.Name, server.GameAddr, server.GamePort, server.LogDBConfig, server.GameData, server.JsonData, server.Status, server.CreateTime, server.LastTime)

		// 解析时间字符串为时间类型
		layout := "2006-1-2 15:04:05"
		ct, err := time.Parse(layout, server.CreateTime)
		if err != nil {
			fmt.Println("解析时间字符串时出错:", err)
			return
		}
		// t := ct.Format(time.RFC3339)

		lt, err := time.Parse(layout, server.LastTime)
		if err != nil {
			fmt.Println("解析时间字符串时出错:", err)
			return
		}

		ds := DbServer{
			SID:          server.ID,
			Name:        server.Name,
			GameAddr:    server.GameAddr,
			GamePort:    server.GamePort,
			LogDBConfig: server.LogDBConfig,
			GameData:    server.GameData,
			JsonData:    server.JsonData,
			Status:      server.Status,
			CreateTime:  ct,
			LastTime:    lt,
		}

		fmt.Println(ds)
		// 插入文档
		result, err := collection.InsertOne(ctx, ds)
		if err != nil {
			log.Fatal(err)
		}

		// 获取插入的文档ID
		insertedID := result.InsertedID
		fmt.Println("Inserted ID:", insertedID)

	}

	fmt.Println("Data inserted successfully!")

}
