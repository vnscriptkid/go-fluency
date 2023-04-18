package lib

import (
	"log"

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

	// q, err := ch.QueueDeclare(
	// 	"task_queue",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// failOnError(err, "Failed to declare a queue")

	// // Consumer
	// go func() {
	// 	msgs, err := ch.Consume(
	// 		q.Name,
	// 		"",
	// 		false,
	// 		false,
	// 		false,
	// 		false,
	// 		nil,
	// 	)
	// 	failOnError(err, "Failed to register a consumer")

	// 	for msg := range msgs {
	// 		log.Printf("Received a message: %s", msg.Body)
	// 		time.Sleep(1 * time.Second)
	// 		msg.Ack(false)
	// 	}
	// }()

	// // Publisher
	// for i := 0; i < 10; i++ {
	// 	body := fmt.Sprintf("Message %d", i)
	// 	err = ch.Publish(
	// 		"",
	// 		q.Name,
	// 		false,
	// 		false,
	// 		amqp.Publishing{
	// 			DeliveryMode: amqp.Persistent,
	// 			ContentType:  "text/plain",
	// 			Body:         []byte(body),
	// 		})
	// 	failOnError(err, "Failed to publish a message")
	// 	log.Printf("Sent a message: %s", body)
	// 	time.Sleep(1 * time.Second)
	// }
}
