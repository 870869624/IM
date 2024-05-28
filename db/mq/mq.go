package db

import (
	"context"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to connect to RabbitMQ")
	que, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to connect to RabbitMQ")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "你好呀"
	err = ch.Publish(
		"",
		que.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to connect to RabbitMQ")
}

func failOnError(err error, str string) {
	if err != nil {
		log.Println(str)
	}
}
