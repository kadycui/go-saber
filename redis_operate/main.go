package main

import (
	"gostart/redis_operate/utils"
	"time"
)

func main() {
	// utils.RedisSet("age", "23", time.Hour)
	utils.Set("class", "艺术设计", time.Hour*24)

}
