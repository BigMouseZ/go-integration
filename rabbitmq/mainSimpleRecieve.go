package main

import (
	"go-integration/rabbitmq/rabbitmqdemo"
)

func main() {
	rabbitmq := rabbitmqdemo.NewRabbitMQSimple("testSimple")
	rabbitmq.ConsumeSimple()
}