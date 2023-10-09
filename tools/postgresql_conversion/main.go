package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kadycui/go-saber/tools/postgresql_conversion/model"
	"github.com/kadycui/go-saber/tools/postgresql_conversion/utils"
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

func main() {
	db := utils.GetDB()
	// 读取JSON数据
	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "./conf/servers.json")
	file, err := os.ReadFile(p2)
	if err != nil {
		log.Fatal(err)
	}

	var servers []model.Server
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

		ds := utils.DbServer{
			ID:          server.ID,
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

		result := db.Create(&ds)
		if result.Error != nil {
			log.Println(result.Error)
		}

	}
	fmt.Println("数据初始化完成!!")

}
