package utils

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func ReadRedisConf() {
	// 读取YAML配置文件
	// viper.SetConfigFile("D:\\GoCode\\github.com/kadycui/go-saber\\redis_operate\\utils\\config.yml")
	viper.SetConfigFile("tools\\redis_operate\\conf\\config.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	// 获取Redis连接信息
	redisHost := viper.GetString("redis.host")
	redisPort := viper.GetString("redis.port")
	redisPassword := viper.GetString("redis.password")

	fmt.Println(redisHost, redisPort, redisPassword)

	// 连接到Redis数据库
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPassword,
	})

	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong) // Output: PONG
}
