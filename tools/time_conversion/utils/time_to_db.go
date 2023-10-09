package utils

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type GameServer struct {
	ID         int
	ClientVer  string
	Name       string
	GameAddr   string
	GamePort   int
	LastTime   time.Time
	CreateTime time.Time
}

func TimeToDb() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	defer db.Close()

	// 创建表
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS game_servers (id INT, client_ver VARCHAR(255), name VARCHAR(255), game_addr VARCHAR(255), game_port INT, last_time DATETIME, create_time DATETIME)")
	if err != nil {
		panic("创建表失败：" + err.Error())
	}

	// 创建GameServer对象
	gameServer := GameServer{
		ID:         5,
		ClientVer:  "v1.3",
		Name:       "2服",
		GameAddr:   "sghj-game-nm1.shyouai.com",
		GamePort:   12003,
		LastTime:   time.Now(), // 假设当前时间为LastTime
		CreateTime: time.Now(), // 假设当前时间为CreateTime
	}

	// 将时间格式化为MySQL支持的日期时间字符串
	lastTimeStr := gameServer.LastTime.Format("2006-01-02 15:04:05")
	createTimeStr := gameServer.CreateTime.Format("2006-01-02 15:04:05")

	// 插入数据
	_, err = db.Exec("INSERT INTO game_servers (id, client_ver, name, game_addr, game_port, last_time, create_time) VALUES (?, ?, ?, ?, ?, ?, ?)", gameServer.ID, gameServer.ClientVer, gameServer.Name, gameServer.GameAddr, gameServer.GamePort, lastTimeStr, createTimeStr)
	if err != nil {
		panic("插入数据失败：" + err.Error())
	}

	fmt.Println("数据插入成功！")
}
