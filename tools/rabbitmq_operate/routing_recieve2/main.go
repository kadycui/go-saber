package main

import rabbitmq3 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_routing"

func main() {
	gtwo := rabbitmq3.NewRabbitMQRouting("goqueue3", "gtwo")
	gtwo.RecieveRouting()
}
