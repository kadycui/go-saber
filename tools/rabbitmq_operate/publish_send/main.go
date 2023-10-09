package main

import (
	"fmt"
	"strconv"
	"time"

	rabbitmq2 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_publish"
)

// Publish模式（订阅模式，消息被路由投递给多个队列，一个消息被多个消费者获取）
// 相关场景:邮件群发,群聊天,广播(广告)
func main() {
	rmq := rabbitmq2.NewRabbitMQPubSub("" + "go2queue")

	for i := 0; i < 100; i++ {
		rmq.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条" + "数据")
		fmt.Println("订阅模式生产第" + strconv.Itoa(i) + "条" + "数据")
		time.Sleep(1 * time.Second)
	}

}
