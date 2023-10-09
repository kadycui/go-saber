package utils

import "github.com/go-redis/redis/v8"

var Db *redis.Client

func init() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.166.56:6379",
		Password: "",
		DB:       0,
	})

	Db = rdb

}

func GetDb() *redis.Client {
	return Db
}
