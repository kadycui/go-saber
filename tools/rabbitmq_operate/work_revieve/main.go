package main

import rabbitmq "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_work"

func main() {
	rmq := rabbitmq.NewRabbitMQSimple("" + "goqueue")
	rmq.ConsumeSimple()
}
