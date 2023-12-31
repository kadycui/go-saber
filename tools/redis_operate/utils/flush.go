package utils

import (
	"fmt"

	"github.com/go-redis/redis"
)

func FlushData() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "172.20.166.56:6379",
		Password: "", // Redis 无密码设置
		DB:       3,        // Redis 默认数据库
	})

	// 调用 FlushDB 清空 Redis 数据库
	_, err := client.FlushDB().Result()

	if err != nil {
		fmt.Println("Failed to flush database:", err)
	} else {
		fmt.Println("Database is flushed.")
	}
}
