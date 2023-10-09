package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/kadycui/go-saber/tools/time_conversion/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbServer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	GameAddr    string    `json:"game_addr"`
	GamePort    int64     `json:"game_port"`
	LogDBConfig string    `json:"log_db_config"`
	GameData    string    `json:"game_data"`
	JsonData    string    `json:"json_data"`
	Status      int64     `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	LastTime    time.Time `json:"last_time"`
}

func GetPath() string {
	// 获取当前执行文件绝对路径（go run）
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get current file path")
	}
	dir := path.Dir(filename)
	return dir
}

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// TableName 指定动态表名称为 dynamic
func (DbServer) TableName() string {
	return "serverS"
}

// Migrate 创建动态表
func Migrate(db *gorm.DB) error {
	// err := db.Table("dynamic").AutoMigrate(&Dynamic{})
	err := db.AutoMigrate(&DbServer{})
	return err
}

func Server2Db() {
	db, err := InitDB()
	if err != nil {
		// 处理初始化数据库失败的情况
		log.Fatal(err)
	} else {
		log.Println("DB初始化完成!")
	}

	err = Migrate(db)
	if err != nil {
		// 处理创建表失败的情况
		log.Fatal(err)
	} else {
		log.Println("数据库建表完成!")
	}

	// 读取JSON数据
	path := GetPath()
	p1, _ := filepath.Abs(path)
	p2 := filepath.Join(p1, "../conf/servers.json")
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

		ds := DbServer{
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

	fmt.Println("Data inserted successfully!")
}
