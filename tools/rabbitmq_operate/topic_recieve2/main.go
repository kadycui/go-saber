package main

import rabbitmq4 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_topic"

func main() {
	gone := rabbitmq4.NewRabbitMQTopic("gotopic", "go.*.two")
	gone.RecieveTopic()
}
