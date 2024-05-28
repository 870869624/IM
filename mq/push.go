package mq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// 发送消息往通道发送消息
func PushMessage(data []byte, queName string) error {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	// a := ch.Confirm(false) //回滚
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
		return err
	}
	defer ch.Close()

	// confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1)) //处理确认逻辑
	// defer confirmOne(confirms)

	ch.Publish(
		"",
		que.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		},
	)
	return nil
}

//确认消息是否投递成功
// func confirmOne(confirms <-chan amqp.Confirmation) {
// 	if confirmed := <-confirms; confirmed.Ack {

// 	}
// }
