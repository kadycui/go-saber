package utils

import (
	"fmt"
	"time"
)

/*
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。引用类型channel是CSP模
式的具体体现，用于多个goroutine之间的通信。其内部实现了同步，确保并发安全

var 通道变量 chan 通道类型


定义一个channel时，也需要定义发送到channel的值的类型。channel可以使用内置的make()函
数来创建：
make(chan Type) // 等价于make(chan Type, 0)
make(chan Type, capacity)

channel通过操作符“＜-”来接收和发送数据，接收和发送数据的语法如下
	channel <- value //发送value到channel
	<-channel //接收并将其丢弃
	x := <-channel //从channel中接收数据，并赋值给x
	x, ok := <-channel //同上，并检查通道是否关闭，将此状态赋值给ok
*/

func Demo7() {
	// 定义并创建了一个可以传输string类型的ch通道变量，在匿名协程函数中，
	//从ch通道中接收数据并打印
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()
	ch <- "test"

	time.Sleep(time.Second)
}

// 缓冲机制
/*
无缓冲的channel创建格式

make(chan Type) //等价于make(chan Type, 0)
*/

func Demo8() {
	ch := make(chan int, 0)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("len(ch)=%v, cap(ch)=%v\n", len(ch), cap(ch))
			ch <- i
		}
	}()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(<-ch)
	}
}

// 有缓冲通道
/*
有缓冲通道是一种在被接收前能存储一个或多个值的通道。创建一个有缓冲通道的方式
make(chan Type, capacity)
*/

func Demo9() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("len(ch)=%v, cap(ch)=%v\n", len(ch), cap(ch))
			ch <- i
		}
	}()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 3)
		fmt.Println(<-ch)
	}
}

/*
这导致有缓冲的通道和无缓冲的通道之间有一个很大的不同：无缓冲的通道保证进行发送和
接收的goroutine会在同一时间进行数据交换，有缓冲的通道没有这种保证

一个容量为3的缓冲channel，由于存在缓冲区，在缓冲区未填满的情况下，程序就不会被阻塞执行
*/

// close和range
/*
◇ channel不像文件一样需要经常去关闭，只有当你确实没有任何需要发送的数据时，或者想
要显式地结束range循环之类的，才会去关闭channel。
◇ 关闭channel后，无法向channel再次发送数据，再次发送将会引发panic错误。
◇ 关闭channel后，可以继续从channel接收数据。
◇ 对于nil channel，无论接收还是发送都会被阻塞。
*/

func Demo10() {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("len(ch)=%v, cap(ch)=%v\n", len(ch), cap(ch))
			ch <- i
		}
		// 发送者在发送完数据后，使用了close关闭channel。channel被关闭后，ok的值就会变为false，最终程序结束运行
		close(ch)
	}()

	//for {
	//	if val, ok := <-ch; ok == true {
	//		fmt.Println(val)
	//	} else {
	//		return
	//	}
	//
	//}

	// 简洁写法
	for data := range ch {
		fmt.Println(data)
	}
}

// 单向channel变量的声明
/*
var ch1 chan int // ch1为一个双向通道
var ch2 chan<- int // ch2为一个只能接收的单向通道
var <-chan int // ch3为一个只能发送的单向通道

"chan＜-"表示数据进入通道，只要把数据写入通道，对于调用者而言就是输出。
"＜-chan"表示数据从通道中出来，对于调用者就是得到通道的数据，也就是输入
*/
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}

	// 不关闭通道就会出现协程死锁 fatal error: all goroutines are asleep - deadlock!
	close(out)
}

func consumer(in <-chan int) {
	for val := range in {
		fmt.Println(val)
	}
}

func Demo11() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

}

func Demo12() {
	// 在Go语言标准库的time包中，定时器的实现使用的就是单向channel
	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		fmt.Println("loop")
	}

}

//通过关键字select可以监听channel上的数据流动。select的用法和switch非常相似，由select开始一个新的选择块，每个选择条件由case语句来描述
/*
与switch语句可以选择任何可使用相等比较的条件相比，select有较多的限制，其中最大的限制
就是每个case语句里面必须是一个I/O操作，大致结构如下

select {
	case <-chan1:
	 // 如果chan1成功读到数据，则执行该case语句
	case chan2 <- 1:
	 // 如果成功向chan2写入数据，则执行该case语句
	default:
	 //如果上面的case都没有执行成功，则执行该default语句
	}
*/

func Demo13() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i

		}
	}()

	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)

		default:
			time.Sleep(time.Second)

		}
	}

}

// 超时
//有时候会出现goroutine阻塞的情况，为了避免程序长时间进入阻塞，我们可以使用select来实现阻塞超时机制

func Demo14() {
	ch := make(chan int)
	done := make(chan bool)
	go func() {
		for {
			select {
			case val := <-ch:
				fmt.Println(val)
			case <-time.After(time.Second * 3):
				fmt.Println("已经超时3秒")
				done <- true
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- i

	}
	<-done
	fmt.Println("程序终止!")

}

// 死锁

/*
只有一个goroutine，我们向里面加数据或者存数据的话，都会锁死信道，并且
阻塞当前goroutine，也就是所有的goroutine（其实只有main线程一个）都在等待信道的开放（没人拿走数据的话，信道是不会开放的），
这就产生了死锁。
在非缓冲信道上如果发生了流入无流出，或者流出无流入，都会导致死锁。同样地，使用
select关键字，其中不加任何代码也会产生死锁

*/

func Demo15() {
	ch := make(chan int)
	<-ch // 阻塞main goroutine, 信道ch被锁

}

func Demo16() {
	select {}

}
