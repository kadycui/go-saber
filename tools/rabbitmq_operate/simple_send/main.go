package main

import (
	"fmt"

	rabbitmq "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_work"
)

func main() {
	rmq := rabbitmq.NewRabbitMQSimple("" + "goqueue")
	rmq.PublishSimple("hello go go go")
	fmt.Println("发送成功!")

}
