package utils

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch = make(chan int, 6)

/*
(1)写入的次数与读取的次数需要一致（本例是10）；
(2)如果读的次数多于写的次数会发生：fatal error: all goroutines are asleep - deadlock! ，若  在mm1中对ch8进行关闭（执行  close(ch8) ），多于的次数读到的数据为0（数据默认值）
(3)读的次数少于写的次数，会读取出次数对应的内容，不会报错。
*/

func ChIn() {
	for i := 0; i < 10; i++ {
		ch <- 8 * i
	}

	fmt.Printf("ch: len= %v cap=%v \n", len(ch), cap(ch))
	close(ch)
}

func ChInWg() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		ch <- 8 * i
	}

	fmt.Printf("ch: len= %v cap=%v \n", len(ch), cap(ch))
	close(ch)
}

func ChannelWg() {
	wg.Add(1)
	go ChInWg()

	for i := 0; i < 100; i++ {
		// 阻塞接收
		fmt.Print(<-ch, "\t")

	}

	wg.Wait()

}

func AssertChannel() {
	go ChIn()

	for {
		// 非阻塞接收
		if data, ok := <-ch; ok {
			fmt.Print(data, "\t")
		} else {
			break
		}
	}

}

func ChannelRange() {
	go ChIn()

	for {
		for data := range ch {
			fmt.Print(data, "\t")

		}
	}
}
