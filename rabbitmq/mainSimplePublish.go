package main

import (
	"fmt"

	"go-integration/rabbitmq/rabbitmqdemo"
)

func main() {
	rabbitmq := rabbitmqdemo.NewRabbitMQSimple("testSimple")
	rabbitmq.PublishSimple("Hello Test! 测试发送MQ")
	fmt.Println("发送成功！")
}

