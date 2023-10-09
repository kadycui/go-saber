package main

import (
	"fmt"
	"strconv"
	"time"

	rabbitmq4 "github.com/kadycui/go-saber/tools/rabbitmq_operate/rabbitmq_topic"
)

// Topic模式（话题模式，一个消息被多个消费者获取，消息的目标queue可用BindingKey以通配符，
// （#：一个或多个词，*：一个词）的方式指定

/*
星号井号代表通配符
星号代表多个单词,井号代表一个单词
路由功能添加模糊匹配
消息产生者产生消息,把消息交给交换机
交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费
*/

func main() {
	gone := rabbitmq4.NewRabbitMQTopic("gotopic", "go.topic.one")
	gtwo := rabbitmq4.NewRabbitMQTopic("gotopic", "go.topic.two")
	for i := 0; i <= 100; i++ {
		gone.PublishTopic("Hello go topic one!" + strconv.Itoa(i))
		gtwo.PublishTopic("Hello go topic Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
