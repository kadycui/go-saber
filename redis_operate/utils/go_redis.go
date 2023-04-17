package utils

// import (
// 	"fmt"
// 	"time"

// 	"github.com/go-redis/redis"
// )

// var RedisDb *redis.Client

// func init() {
// 	// 初始化redis
// 	RedisDb = redis.NewClient(&redis.Options{
// 		Addr: "127.0.0.1:6379",
// 	})

// 	_, err := RedisDb.Ping().Result()
// 	if err != nil {
// 		fmt.Println("链接redis失败！", err)
// 		return

// 	}

// 	fmt.Println("链接成功！！")
// }

// func RedisSet(key string, value interface{}, expiration time.Duration) (interface{}, error) {
// 	err := RedisDb.Set(key, value, expiration).Err()
// 	if err != nil {
// 		return fmt.Sprintln("Redis设置指失败"), err
// 	}
// 	return fmt.Sprintln("设置成功"), nil
// }

// func RedisGet(key string) (interface{}, error) {
// 	value, err := RedisDb.Get(key).Result()
// 	if err != nil {
// 		fmt.Println("Redis获取key失败")
// 		return nil, err
// 	}
// 	return value, nil
// }
