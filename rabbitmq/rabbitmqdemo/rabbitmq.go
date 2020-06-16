package rabbitmqdemo

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// 创建连接url
const MQURL = "amqp://admin:admin@192.168.147.129:5672/"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel

	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// key
	Key string
	// 连接信息
	Mqurl string
}

// 创建连接实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	// exchange 为空会使用默认的default
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
	var err error
	// 创建rabbitmq来连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误！")

	// 创建channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败")
	return rabbitmq
}

// 断开连接:channel和conn
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// 错误处理的函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
		panic(fmt.Sprintf("%s", message))
	}
}

// step1: simple style 创建简单模式的实例, 只需要队列名
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// step2: 简单模式下生产
func (r *RabbitMQ) PublishSimple(message string) {
	// 固定用法  申请队列，如果队列不存在会自动创建，如果存在则直接使用，保证队列中能存入数据
	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名
		false,       // 控制是否持久化
		false,       // 是否自动删除，当最后一个消费者断开连接后是否删除
		false,       // 是否具有排他性，其他用户不可访问
		false,       // 是否阻塞
		nil,         //  额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false, // mandatory 如果为true,会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把消息返回给发送者
		false, // immediate 如果为true,当exchange发送消息到队列后发现队列没有绑定消费者后会把消息返回
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

// 简单模式的消费消息
func (r *RabbitMQ) ConsumeSimple() {
	// 固定用法  申请队列，如果队列不存在会自动创建，如果存在则直接使用，保证队列中能存入数据
	_, err := r.channel.QueueDeclare(
		r.QueueName, // 队列名
		false,       // 控制是否持久化
		false,       // 是否自动删除，当最后一个消费者断开连接后是否删除
		false,       // 是否具有排他性，其他用户不可访问
		false,       // 是否阻塞
		nil,         //  额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	// 接受消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",     // 用来区分多个消费在
		true,   // 是否自动应答， 主动的告诉mq自己已经消费完了，如果false,需要回调函数
		false,  // 是否排他性
		false,  // 如果设置为true, 表示不能将同一个connection中发送的消息传递给这个connect中的消费者
		false,  // 设置为阻塞，一个一个消费
		nil,    // 附加信息
	)
	if err != nil {
		fmt.Println(err)
	}

	// 消费时的固定写法,用来阻塞
	forever := make(chan bool)
	// 启用协程处理消息
	go func() {
		for d := range msgs {
			// 实现我们要处理的逻辑函数
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for message,To exit press CTRL + C")
	<- forever
}