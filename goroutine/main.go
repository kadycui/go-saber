package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()  //计数-1
	fmt.Println("Hello Goroutine!")
}

func word() {
	defer wg.Done()
	fmt.Println("Word Goroutine!")
}

func main() {
	wg.Add(2) // 计数器+1

	go hello()

	go word()

	wg.Wait()  // 阻塞到计数器变为0

	fmt.Println("Main Goroutine!")

	// time.Sleep(2 * time.Second)

}
