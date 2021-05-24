package rabbitmq_consumer

import (
	"fmt"
	"github.com/streadway/amqp"
	// local packages
  "github.com/gpenaud/needys-output-producer/internal/config"
	"github.com/gpenaud/needys-output-producer/pkg/log"
)

func WaitForAmqpMessages() {
	rabbitmq_connection_parameters := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.Cfg.Rabbitmq.Username,
		config.Cfg.Rabbitmq.Password,
		config.Cfg.Rabbitmq.Host,
		config.Cfg.Rabbitmq.Port,
	)

	conn, err := amqp.Dial(rabbitmq_connection_parameters)
	if err != nil {
		log.ErrorLogger.Fatalln("Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.ErrorLogger.Fatalln("Failed to open a channel")
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
		log.ErrorLogger.Fatalln("Failed to declare a queue")
	}

	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.ErrorLogger.Fatalln("Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for m := range messages {
			log.InfoLogger.Printf("Received a message: %s", m.Body)
		}
	}()

	log.InfoLogger.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
