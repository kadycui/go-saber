package utils

import "fmt"

var intChan chan int

func ChUse() {
	//演示一下管道的使用
	//1. 创建一个可以存放3个int类型的管道

	intChan = make(chan int, 3)

	//2. 看看intChan是什么
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	//3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	// //如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98 //注意点, 当我们给管写入数据时，不能超过其容量

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据

	num2 := <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock

	num3 := <-intChan
	num4 := <-intChan

	//num5 := <-intChan

	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

}

type Cat struct {
	Name string
	Age  int64
}

func EmptyInterfaceChan() {
	// 定义一个存放任意数据类型的管道  3个数据
	// var allChan chan interface{}

	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "Jack Ma"
	cat := Cat{"波斯猫", 6}
	allChan <- cat

	// 要获取管道中的第三个元素, 需要把前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)

	//下面的写法是错误的!编译不通过
	//fmt.Printf("newCat.Name=%v", newCat.Name)

	// 定义 interface类型的空接口，可以接收任意类型的数据，但是在取出来的时候，必须断言！
	//a := newCat.(Cat)

	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v \n", a.Name)

}

func ConcurrentSync() {
	// 通过通道在 goroutine 间阻塞收发实现并发同步
	ch := make(chan int)

	go func() {
		fmt.Println("开启goroutine")

		// 通过通道通知当前主函数的goroutine
		ch <- 0

		fmt.Println("退出goroutine")

	}()

	fmt.Println("等待goroutine")

	// 等待匿名goroutine
	<-ch

	fmt.Println("运行结束")

}

func CloseChannel() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	close(intChan) // 关闭管道

	//这时不能够再写入数到channel
	// intChan <- 300

	fmt.Println("okook~")

	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)

	// 遍历管道
	intChan2 := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道

	}

	//遍历管道不能使用普通的 for 循环
	// for i := 0; i < len(intChan2); i++ {

	// }
	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，就会退出遍历

	close(intChan2)   // 在遍历管道之前要先关闭管道，不然会出现deadlock的错误

	for v := range intChan2 {
		fmt.Print(v, "\t")
	}

}
