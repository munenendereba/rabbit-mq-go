package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"

	"munendereba/rabbit-mq-service/helper"
)

func Send() {
	conn := helper.ConnectRabbitMQServer()

	ch, err := conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")

	defer ch.Close()
	defer conn.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	helper.FailOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	messages := []string{"muuga", "hello people", "how can i help you", "quick sand is not good"}

	for _, msg := range messages {
		body := msg
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		helper.FailOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", body)
	}
}
