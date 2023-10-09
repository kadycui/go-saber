package rabbitmq3

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// 连接信息
const MQURL = "amqp://guest:guest@172.20.166.56:5672"

// rabbitmq 结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	// 队列名称
	QueueName string

	// 交换机名
	Exchange string

	// bind Key 名称
	Key string

	// 连接信息
	Mqurl string
}

// 创建结构体实列
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// 断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理函数
func (r *RabbitMQ) failOnErr(err error, messge string) {
	if err != nil {
		log.Fatal("%s:%s", messge, err)
		panic(fmt.Sprintf("%s:%s", messge, err))
	}

}

// 路由模式
// 创建RabbitMQ实列
func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	// 创建RabbitMQ实列
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)

	var err error

	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "rabbitmq 连接失败!")

	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "打开channel失败")

	return rabbitmq

}

// 路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) {

	// 1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,

		// 要改成direct
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an excha"+"nge")

	// 2.发送消息
	err = r.channel.Publish(
		r.Exchange,

		// 设置Key
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

// 路由模式接受消息
func (r *RabbitMQ) RecieveRouting() {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+"ange")

	// 2.尝试创建队列
	q, err := r.channel.QueueDeclare(
		"", // 随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	// 绑定队列到exchange中
	err = r.channel.QueueBind(
		q.Name,
		// 需要绑定key
		r.Key,
		r.Exchange,
		false,
		nil)
	// 消费消息
	message, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	forever := make(chan bool)

	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}

	}()
	fmt.Println("退出请按 CTRL+C\n")
	<-forever

}
