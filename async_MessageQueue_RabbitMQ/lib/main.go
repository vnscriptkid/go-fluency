package lib

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Exec() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Println("Connected to RabbitMQ")

	q, err := ch.QueueDeclare(
		"task_queue",
		true,  // Durable: true (the queue will survive a broker restart)
		false, // Auto-delete: false (the queue will not be deleted when there are no consumers)
		false, // Exclusive: false (the queue can be accessed by other connections)
		false, // No-wait: false (the queue declaration will be synchronous)
		nil,   // Arguments: nil (no additional arguments)
	)
	failOnError(err, "Failed to declare a queue")

	// Consumer
	go func() {
		msgs, err := ch.Consume(
			q.Name, // Queue: q.Name (the name of the queue)
			"",     // Consumer: "" (empty string means the server will generate a unique consumer tag)
			false,  // Auto-ack: false (the consumer will send manual acknowledgements after processing messages)
			false,  // Exclusive: false (the queue is not exclusive to this consumer)
			false,  // No-local: false (the consumer can receive messages published on the same connection)
			false,  // No-wait: false (the consumer registration will be synchronous)
			nil,    // Arguments: nil (no additional arguments)
		)
		failOnError(err, "Failed to register a consumer")

		for msg := range msgs {
			log.Printf("-- Received message: %s", msg.Body)
			log.Printf("@@ Processing message: %s", msg.Body)
			time.Sleep(1 * time.Second) // Simulate processing time with a 1-second sleep.
			log.Printf("++ Processed message: %s", msg.Body)

			// Send an acknowledgement for the received message.
			// false argument indicates that we're sending an acknowledgement for just this single message,
			// rather than multiple messages at once
			// The message will not be automatically deleted simply because it has not been acknowledged
			if fmt.Sprintf("%s", msg.Body) == "Message 7" {
				log.Printf("!! Simulating msg not being acked: %s", msg.Body)
			} else {
				msg.Ack(false)
				log.Printf("++ Processed message: %s", msg.Body)
			}
		}
	}()

	// Publisher
	for i := 7; i <= 7; i++ {
		body := fmt.Sprintf("Message %d", i)
		err = ch.Publish(
			"",     // Exchange: "" (empty string means the default exchange)
			q.Name, // Routing key: q.Name (the name of the queue)
			false,  // Mandatory: false (the message will be discarded if it can't be routed)
			false,  // Immediate: false (the message can be buffered if the consumer is busy)
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // the message will be stored to disk and survive broker restarts
				ContentType:  "text/plain",    // the message body is plain text
				Body:         []byte(body),    // the message body
			})
		failOnError(err, "Failed to publish a message")
		log.Printf("Sent a message: %s", body)
		time.Sleep(1 * time.Second)
	}

	forever := make(chan bool)
	forever <- true
}
