package utils

import (
	"context"
	"fmt"
)

func Demo1() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	c2 <- "hello"

	select {
	case msg1 := <-c1:
		fmt.Println("c1的返回值: ", msg1)
	case msg2 := <-c2:
		fmt.Println("c2的返回值: ", msg2)

	default:
		fmt.Println("没有数据返回!")

	}
}

func Demo2() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 结束 goroutine 运行
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 消耗完时取消

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
