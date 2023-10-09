package main

import (
	"strconv"
	"time"

	rabbitmq "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_work"
)

//Work模式（工作模式，一个消息只能被一个消费者获取）

func main() {
	rmq := rabbitmq.NewRabbitMQSimple("" + "goqueue")
	for i := 0; i <= 100; i++ {
		rmq.PublishSimple("hello go go --" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
