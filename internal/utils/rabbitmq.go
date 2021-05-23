package utils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	// local packages
  "github.com/gpenaud/needys-output-producer/internal/models"
)

func WaitForAmqpMessages() {
	rabbitmq_connection_parameters := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		models.Cfg.Rabbitmq.Username,
		models.Cfg.Rabbitmq.Password,
		models.Cfg.Rabbitmq.Host,
		models.Cfg.Rabbitmq.Port,
	)

	conn, err := amqp.Dial(rabbitmq_connection_parameters)
  FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for m := range messages {
			log.Printf("Received a message: %s", m.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
