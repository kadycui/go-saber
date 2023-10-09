package main

import rabbitmq2 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_publish"

func main() {
	rmq := rabbitmq2.NewRabbitMQPubSub("" + "go2queue")
	rmq.RecieveSub()
}
