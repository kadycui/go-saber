package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
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

// func initDb() redis.Client {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:         "172.20.166.56:6379",
// 		Password:     "",
// 		DB:           0,
// 		DialTimeout:  10 * time.Second,
// 		ReadTimeout:  30 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 		PoolSize:     10,
// 		PoolTimeout:  30 * time.Second,
// 	})

// 	return *rdb
// }

func JsonToRedis() {
	// 读取JSON数据
	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "../conf/data.json")
	file, err := os.ReadFile(p2)
	if err != nil {
		log.Fatal(err)
	}

	// var players []Player
	// err = json.Unmarshal([]byte(file), &players)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, player := range players {

	// 	// 解析时间字符串为时间类型
	// 	// layout := "2006-1-2 15:04:05"

	// 	fmt.Println(player)

	// }

	var result []map[string]interface{}

	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println("无法解析JSON数据:", err)
		return
	}

	rdb := NewRedisClient()

	defer rdb.Close()

	for _, p := range result {
		intValue, err := strconv.ParseInt(fmt.Sprintf("%.0f", p["player_id"]), 10, 64)
		if err == nil {
			// 将转换后的整数替换回map中的值
			p["player_id"] = intValue
		}

		index := p["player_id"]

		// 将interface{}类型转换为int64类型
		intValue, ok := index.(int64)
		if !ok {
			fmt.Println("无法将interface{}转换为int64类型")
			return
		}

		// 将int64类型转换为string类型
		str := strconv.FormatInt(intValue, 10)

		data, err := json.Marshal(p)
		if err != nil {
			panic(err)
		}

		jsonData := string(data)
		fmt.Println(jsonData)

		// err = rdb.Set(ctx, str, jsonData, 0).Err()
		// if err != nil {
		// 	panic(err)
		// }

		err = rdb.Set(str, jsonData, 0)
		if err != nil {
			fmt.Println("添加数据失败!!")
		}

	}

}
