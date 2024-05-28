package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// 消费通道消息
func ConsumMessage(queName string) *<-chan amqp.Delivery {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer ch.Close()

	que, err := ch.QueueDeclare(
		queName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer ch.Close()

	//不应该自动返回ack，如果消费为成功则会出现bug
	msgs, err := ch.Consume(
		que.Name,
		"",
		false, false, false, false, nil,
	)
	if err != nil {
		panic("4444" + err.Error())
	}
	// for d := range msgs {
	// 	fmt.Println(string(d.Body))
	// 	d.Ack(false)
	// }
	//消费消息
	return &msgs
}
