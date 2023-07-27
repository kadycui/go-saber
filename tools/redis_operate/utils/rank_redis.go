package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func MapToRedis() {

	type Rank struct {
		Member string
		Score  float64
	}

	var ranks []Rank
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
		Addr:     "10.16.168.61:6379",
		Password: "123456", // Redis 无密码
		DB:       3,        // 数据库编号
	})
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			panic(err)
		}
	}(rdb)

	// 清空db
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

		data := map[string]interface{}{
			"server_id":    f2 >> 20,
			"player_id":    f2,
			"player_name":  f3,
			"player_level": f5,
			"gang":         f4,
			"battle":       battle,
			"world_Level":  f7,
		}

		r := Rank{Member: strconv.Itoa(f2), Score: float64(battle)}
		ranks = append(ranks, r)

		err = rdb.HMSet(fmt.Sprintf("lbr|%d", f2), data).Err()
		if err != nil {
			fmt.Println(err)
		}
	}

	// map数据存入有序集合
	zs := make([]redis.Z, len(ranks))
	for i, r := range ranks {
		zs[i] = redis.Z{Score: r.Score, Member: r.Member}
	}

	if err := rdb.ZAdd("lbr|rank", zs...).Err(); err != nil {
		fmt.Println(err)
	}

	// 使用 ZRevRange 命令获取指定排名范围内的元素
	result, err := rdb.ZRevRange("lbr|rank", 0, 10).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("ZRevRange result: %+v", result)

	var ps []map[string]interface{}
	// 遍历获取到的结果
	for _, member := range result {
		val, err := rdb.HGetAll(fmt.Sprintf("lbr|%s", member)).Result()
		if err != nil {
			fmt.Println(err)
		} else {
			mapInterface := make(map[string]interface{})
			for k, v := range val {
				mapInterface[k] = v
			}
			ps = append(ps, mapInterface)
		}
	}

	jsonBytes, err := json.Marshal(ps)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

}
