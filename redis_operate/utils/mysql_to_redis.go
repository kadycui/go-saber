package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlToRedis() {

	type Player struct {
		PlayerId   int    `json:"player_id"`
		PlayerName string `json:"player_name"`
		Gang       string `json:"gang"`
		BattleNum  int    `json:"battle_num"`
	}

	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 查询需要的数据
	rows, err := db.Query("SELECT f2, f3, f4, cast(f6 as UNSIGNED)  as f6_num FROM log_battle_rank GROUP BY f2 order by f6_num  desc  limit 100;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 连接 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.16.168.61:6379",
		Password: "123456", // Redis 无密码
		DB:       3,        // 数据库编号
	})
	defer rdb.Close()

	// 存储数据到 Redis
	for rows.Next() {
		var f2 int
		var f3 string
		var f4 string
		var f6_num int
		err := rows.Scan(&f2, &f3, &f4, &f6_num)
		if err != nil {
			panic(err)
		}

		p := Player{
			PlayerId:   f2,
			PlayerName: f3,
			Gang:       f4,
			BattleNum:  f6_num,
		}

		// 将用户编码为JSON格式
		player_info, err := json.Marshal(p)
		if err == nil {
			fmt.Println(string(player_info))
		} else {
			fmt.Println(err)
			return
		}

		// 将JSON格式的数据存储到Redis
		// err = rdb.Set(fmt.Sprintf("p:%d", f2), player_info, 0).Err()

		err = rdb.HSet(fmt.Sprintf("p:%d", f2), fmt.Sprint(f2), string(player_info)).Err()
		if err != nil {
			panic(err)
		}

		// 从Redis中读取JSON格式的数据
		// val, err := rdb.Get(fmt.Sprintf("p:%d", f2)).Result()

		val, err := rdb.HGet(fmt.Sprintf("p:%d", f2), fmt.Sprint(f2)).Result()
		if err != nil {
			panic(err)
		}

		// 解码JSON数据
		var p2 Player
		err = json.Unmarshal([]byte(val), &p2)
		if err != nil {
			panic(err)
		}
		// 输出用户信息
		fmt.Println(p.PlayerId, p.PlayerName, p.Gang, p.BattleNum)
	}

	// 检查是否有错误
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
