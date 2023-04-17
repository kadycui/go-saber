package utils

//
//import (
//	"fmt"
//	"sync"
//)
//
//var lock sync.RWMutex
//
//func readMap(Gomap map[int]int, key int) int {
//	lock.Lock()
//	m := Gomap[key]
//	lock.Unlock()
//	return m
//}
//
//func writeMap(Gomap map[int]int, key int, value int) {
//	lock.Lock()
//	Gomap[key] = value
//	lock.Unlock()
//}
//
//func Demo1() {
//	GoMap := make(map[int]int)
//	for i := 0; i < 10000; i++ {
//		go writeMap(GoMap, i, i)
//		go readMap(GoMap, i)
//	}
//	fmt.Println("Done")
//}
