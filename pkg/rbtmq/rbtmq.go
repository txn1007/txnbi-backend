package rbtmq

import amqp "github.com/rabbitmq/amqp091-go"

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	//exchang, err := ch.ExchangeDeclare()
}
