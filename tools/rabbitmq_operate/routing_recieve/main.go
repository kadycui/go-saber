package main

import rabbitmq3 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_routing"

func main() {
	gone := rabbitmq3.NewRabbitMQRouting("goqueue3", "gone")
	gone.RecieveRouting()
}
