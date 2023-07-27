package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

// sync.WaitGroup

var wg sync.WaitGroup

// 定义计数器
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int64
}

// 计数器生成
func (s *SafeCounter) Inc(key string) {
	defer wg.Done()

	s.mu.Lock() // 加锁

	s.v[key]++ // 对计数器值 增值+1

	s.mu.Unlock() // 解锁
}

// 获取计数器
func (s *SafeCounter) Value(key string) int64 {
	defer wg.Done()

	// 加锁
	s.mu.Lock()

	d := s.v[key]

	// 解锁
	s.mu.Unlock()

	return d
}

func RunCounter() {
	// 初始化计数器
	c := SafeCounter{v: make(map[string]int64)}

	// 开启1000个数生成计数器
	for i := 0; i < 1000; i++ {
		wg.Add(i)
		go c.Inc("aaa")
	}

	fmt.Println("打印计数器:", c.Value("aaa"))

}

// sync.Once

func SyncOnce() {
	o := &sync.Once{}

	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("只执行一次!!")
		})
	}
}

// sync.Cond

var status int64

func broadcast(c *sync.Cond) {
	c.L.Lock()

	atomic.StoreInt64(&status, 1)

	c.Broadcast()
	c.L.Unlock()

}

func listen(c *sync.Cond) {
	c.L.Lock()

	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}

	fmt.Println("listen")

	c.L.Unlock()
}

func SyncCond() {
	c := sync.NewCond(&sync.Mutex{})

	for i := 0; i < 10; i++ {
		go listen(c)
	}

	time.Sleep(1 * time.Second)

	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func main() {

	RunCounter()
	SyncOnce()

	SyncCond()

}
