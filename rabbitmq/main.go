package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	InitRabbitMQ()
}

func InitRabbitMQ() {
	// 连接RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@192.168.147.129:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatal(err)
	}
	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal(err)
	}
	/*
	   接收
	   连接，创建通道，队列

	   在接收端我们同样需要像发送端一样连接RabbitMQ，创建通道后再创建队列，注意此处队列的创建是跟发送端的队列完全匹配的。队列在接收端也创建是因为我们接收端有可能比发送端先启动，所以为了保证我们要消费的队列存在我们在此处也进行创建


	*/
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan string)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	log.Printf(" [*] Waiting for messages. To exit press CTRL+B")
	 <-forever
	log.Printf(" [*] Waiting for messages. To exit press CTRL+D")
}
