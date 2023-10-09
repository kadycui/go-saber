package main

import (
	"fmt"
	"strconv"
	"time"

	rabbitmq3 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_routing"
)

func main() {
	gone := rabbitmq3.NewRabbitMQRouting("goqueue3", "gone")
	gtwo := rabbitmq3.NewRabbitMQRouting("goqueue3", "gtwo")

	for i := 0; i <= 100; i++ {
		gone.PublishRouting("hello gone" + strconv.Itoa(i))
		gtwo.PublishRouting("hello gtwo" + strconv.Itoa(i))

		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
