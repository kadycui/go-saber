package utils

import (
	"fmt"
	"sync"
)

/*
Go在1.9版本中提供的一种效率较高的并发安全的map——sync.Map
sync.Map有以下特点：
◇ 内部通过冗余的数据结构降低加锁对性能的影响。
◇ 使用前无须初始化，直接声明即可。
◇ sync.Map不使用map中的方式来进行读取和赋值等操作。
*/

func Demo2() {
	var GoMap sync.Map
	for i := 0; i < 100000; i++ {
		go writeMap(GoMap, i, i)
		go readMap(GoMap, i)

	}
	fmt.Println("Done")
}

func readMap(GoMap1 sync.Map, key int) int {
	res, ok := GoMap1.Load(key) // 线程安全读取
	if ok == true {
		return res.(int)
	} else {
		return 0
	}
}

func writeMap(GoMap1 sync.Map, key int, value int) {
	GoMap1.Store(key, value) // 线程安全设置

}

/*
◇ sync.Map无须使用make创建。
◇ Load()方法的第一个返回值是接口类型，需要将其转换为map值的类型。
◇ 目前sync.Map没有提供获取map数量的方法，解决方案是通过循环遍历map。
◇ 与较普通的map相比，sync.Map为了保证并发安全，会有性能上的损失，因此在非并发情况
下，推荐使用map。

*/
