package rank

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func JsonToRedis() {

	type Rank struct {
		Member string
		Score  float64
	}

	var ranks []Rank

	type Player struct {
		ServerId    int    `json:"server_id"`
		PlayerId    int    `json:"player_id"`
		PlayerName  string `json:"player_name"`
		PlayerLevel int    `json:"player_level"`
		Gang        string `json:"gang"`
		Battle      int    `json:"battle"`
		WorldLevel  int    `json:"world_Level"`
	}

	// 连接 MySQL 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// 查询需要的数据
	rows, err := db.Query("SELECT f2, MAX(f3), MAX(f4), MAX(f5), MAX(cast(f6 as UNSIGNED)) as battle, MAX(f7) FROM log_battle_rank GROUP BY f2 ORDER BY battle DESC LIMIT 100;")
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	// 连接 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.166.56:6379",
		Password: "123456", // Redis 无密码
		DB:       3,        // 数据库编号
	})
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	_, err = rdb.FlushDB().Result()
	if err != nil {
		panic(err)

	}

	// 存储数据到 Redis
	for rows.Next() {
		var f2 int
		var f3 string
		var f4 string
		var f5 int
		var battle int
		var f7 int
		err := rows.Scan(&f2, &f3, &f4, &f5, &battle, &f7)
		if err != nil {
			panic(err)
		}

		p := Player{
			ServerId:    f2 >> 20,
			PlayerId:    f2,
			PlayerName:  f3,
			Gang:        f4,
			PlayerLevel: f5,
			Battle:      battle,
			WorldLevel:  f7,
		}

		r := Rank{Member: strconv.Itoa(f2), Score: float64(battle)}
		ranks = append(ranks, r)
		// 将用户编码为JSON格式
		playerInfo, err := json.Marshal(p)
		if err == nil {
			fmt.Println("数据库中的数据:", string(playerInfo))
		} else {
			fmt.Println(err)
			return
		}

		// 将JSON格式的数据存储到Redis
		// err = rdb.Set(fmt.Sprintf("p:%d", f2), player_info, 0).Err()

		err = rdb.HSet(fmt.Sprintf("lbr|%d", f2), fmt.Sprint(f2), string(playerInfo)).Err()
		if err != nil {
			panic(err)
		}

		// 从Redis中读取JSON格式的数据
		// val, err := rdb.Get(fmt.Sprintf("p:%d", f2)).Result()
		//
		//val, err := rdb.HGet(fmt.Sprintf("lbr|%d", f2), fmt.Sprint(f2)).Result()
		//if err != nil {
		//	panic(err)
		//}
		//
		//// 解码JSON数据
		//var p2 Player
		//err = json.Unmarshal([]byte(val), &p2)
		//if err != nil {
		//	panic(err)
		//}
		//// 输出用户信息
		//fmt.Println("Redis取出的数据:", p.PlayerId, p.PlayerName, p.Gang, p.PlayerLevel, p.Battle, p.WorldLevel)
	}

	zs := make([]redis.Z, len(ranks))
	for i, r := range ranks {
		zs[i] = redis.Z{Score: r.Score, Member: r.Member}
	}

	if err := rdb.ZAdd("myzset", zs...).Err(); err != nil {
		fmt.Println(err)
	}

}
