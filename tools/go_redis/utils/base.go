package utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Test() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.20.166.56:6379",
		Password: "",
		DB:       1,
	})

	//返回值是当前列表元素的数量
	n, err := rdb.LPush(ctx, "list1", 1, 2, 3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)

}
