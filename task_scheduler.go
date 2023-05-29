package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Error handler
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// this connects to the rabbitmq server running on localhost
	// go uses := to declare and assign a variable on the right
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// if there is an error, log it and exit
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// here we create a channel and assign it to ch
	ch, err := conn.Channel()
	// if there is an error, log it and exit
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// here we declare a queue named "task_queue"
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	// body will be the message we send to the queue
	body := "This is a task message test."
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			// looks similar to the fetch api in js
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		},
	)
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}
