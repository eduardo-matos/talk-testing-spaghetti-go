package main

import (
	"log"
	"os"

	"github.com/annoying-external-dependency/mailer"
	"github.com/annoying-external-dependency/messaging"
	"github.com/streadway/amqp"
)

type email struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"emails", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for d := range msgs {
		emailMsg, err := mailer.ParseMsg(d)

		if err != nil {
			log.Printf("Could not parse message into e-mail: %s", err.Error())
			log.Printf("Ignoring it...")
			continue
		}

		go messaging.OnMsg(emailMsg, os.Stdout)
	}
}
