package utils

import "fmt"

/*
简单来说，接口赋值存在以下两种情况：
1.将对象实例赋值给接口；
2.将一个接口赋值给另一个接口。

当一个对象的类型是一个接口的实例时，这个对象就可以赋值给这个接口。需要注意的是，
只能将对象的指针赋值给接口变量，不能将对象值直接赋值给接口变量，否则就会发生错误
*/

type IDatabaser interface {
	Connect() error
	Disconnect() error
}

type Redis struct {
	DBName string
}

func (redis *Redis) Connect() error {
	fmt.Println("Redis Connect DB =>" + redis.DBName)
	// do Connect
	fmt.Println("Redis Connect Success!")
	return nil
}

func (redis *Redis) Disconnect() error {
	// do Disconnect
	fmt.Println("Redis Disconnect Success!")
	return nil
}

func Demo3() {
	var redis = Redis{DBName: "teacher"}

	// 1-将对象实例赋值给接口
	var idb IDatabaser = &redis
	idb.Connect()
	idb.Disconnect()
}

type IRediser interface {
	Connect() error
}

func Demo4() {
	// 2-通过一个接口给另一个接口赋值
	var idb IDatabaser = &Redis{DBName: "teacher"}

	var iredis IRediser
	iredis = idb
	iredis.Connect()

}
