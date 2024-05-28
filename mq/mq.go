package mq

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

var word []string

func PushMQ() {
	word = []string{"你", "我", "他", "中国", "红色", "任命", "希望"}
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to ch to RabbitMQ")
	que, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to que to RabbitMQ")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 20; i++ {
		flag := rand.Intn(7)

		body := word[flag]
		err = ch.Publish(
			"",
			que.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Failed to publish to RabbitMQ")
	}

}

func failOnError(err error, str string) {
	if err != nil {
		log.Println(str)
	}
}

// func consume() {
// 	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
// 	failOnError(err, "Failed to connect to RabbitMQ")
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	failOnError(err, "Failed to ch to RabbitMQ")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"hello",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	failOnError(err, "Failed to declare to RabbitMQ")

// 	msgs, err := ch.Consume(
// 		q.Name,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	failOnError(err, "Failed to declare to RabbitMQ")

// 	var fo chan struct{}

// 	go func() {
// 		for d := range msgs {
// 			log.Println(string(d.Body))
// 		}
// 	}()
// 	log.Printf("1111")
// 	<-fo
// }

func ExchangeMq() {
	word = []string{"你", "我", "他", "中国", "红色", "任命", "希望"}

	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	//扇区交换
	que, err := ch.QueueDeclare(
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var input string
	for {
		fmt.Scanln(&input)
		body := []byte(input)

		err = ch.Publish(
			"",
			que.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		if err != nil {
			panic(err)
		}
	}

}

// func bodyForm(arg []string) string {
// 	var s string

// 	if len(arg) < 2 || os.Args[1] == "" {
// 		s = "hello"
// 	} else {
// 		s = strings.Join(arg[1:], " ")
// 	}
// 	return s
// }

// func severityForm(args []string) string {
// 	var s string
// 	if (len(args) < 2) || os.Args[1] == "" {
// 		s = "info"
// 	} else {
// 		s = os.Args[1]
// 	}
// 	return s
// }
