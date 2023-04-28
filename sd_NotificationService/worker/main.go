package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/streadway/amqp"
)

const (
	deadQ    = "dead_letter_queue"
	deadQKey = "dead_letter_key"
	deadExc  = "dead_letter"
)

type Message struct {
	Recipient string `json:"recipient"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Channel   string `json:"channel"`
}

func main() {
	// Supported channels:
	// Channel 1: email
	// Channel 2: sms
	// Channel 3: push
	fmt.Println("Hello, world from worker")

	queue := os.Getenv("QUEUE_NAME")
	if queue == "" {
		log.Fatal("QUEUE_NAME environment variable is not set")
		return
	}

	// Start workers for the specified queue
	for i := 0; i < 5; i++ {
		go startWorker(queue, processMessage)
	}

	// Keep the main Goroutine alive
	select {}
}

func processMessage(msg []byte) error {
	// Process the message, e.g., call an external notification sending service
	machineId := os.Getenv("MACHINE_ID")

	msgJson := &Message{}

	err := json.Unmarshal(msg, msgJson)

	if err != nil {
		return err
	}

	// Validate the recipient
	err = checkRecipient(msgJson.Channel, msgJson.Recipient)

	if err != nil {
		return err
	}

	switch msgJson.Channel {
	case "email":
		// Send email
		fmt.Printf("[%s]...sending email to %s\n", machineId, msgJson.Recipient)
		time.Sleep(1 * time.Second)
		fmt.Printf("[%s] email sent to %s\n", machineId, msgJson.Recipient)
	case "sms":
		// Send SMS
		fmt.Printf("[%s]...sending SMS to %s\n", machineId, msgJson.Recipient)
		time.Sleep(2 * time.Second)
		fmt.Printf("[%s] SMS sent to %s\n", machineId, msgJson.Recipient)
	case "push":
		// Send push notification
		fmt.Printf("[%s]...sending push notification to %s\n", machineId, msgJson.Recipient)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("[%s] push notification sent to %s\n", machineId, msgJson.Recipient)
	default:
		return fmt.Errorf("unknown channel: %s", msgJson.Channel)
	}

	return nil
}

func startWorker(queue string, handleMessage func(msg []byte) error) error {
	fmt.Println("Starting worker...")
	host := os.Getenv("RABBITMQ_HOST")

	if host == "" {
		host = "localhost"
	}

	var conn *amqp.Connection

	for {
		_conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:5672/", host))

		if err != nil {
			fmt.Println("Error connecting to RabbitMQ:", err)
			time.Sleep(1 * time.Second)
			fmt.Println("Retrying...")
			continue
		}

		conn = _conn

		break
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Error opening channel:", err)
		return err
	}
	defer ch.Close()

	err = setupDeadLetterQueue(ch)

	if err != nil {
		return err
	}

	fmt.Println("Dead letter queue set up")
	// Set up main queue
	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    deadExc,
			"x-dead-letter-routing-key": deadQKey,
		},
	)
	if err != nil {
		fmt.Println("Error declaring queue", err)
		return err
	}

	fmt.Println("Queue declared")
	msgs, err := ch.Consume(
		q.Name,
		"",
		false, // Set autoAck to false
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	fmt.Println("Waiting for messages...")
	for msg := range msgs {
		fmt.Println("Received a message:", string(msg.Body))
		err := handleMessage(msg.Body)
		if err != nil {
			log.Printf("Error handling message from %s: %v", queue, err)
			err = ch.Publish(
				deadExc,  // Use the dead_letter exchange
				deadQKey, // Use the same routing key as the one used when binding the queue
				false,
				false,
				amqp.Publishing{
					ContentType:  "application/json",
					Body:         msg.Body,
					DeliveryMode: amqp.Persistent,
				},
			)

			if err != nil {
				log.Printf("Error publishing to dead letter queue: %v", err)
			}
		}

		// Acknowledge the message
		err = msg.Ack(false)

		if err != nil {
			log.Printf("Error acknowledging message: %v", err)
		}
	}

	return nil
}

func setupDeadLetterQueue(ch *amqp.Channel) error {
	// Declare a direct exchange for dead-letter queues:
	err := ch.ExchangeDeclare(deadExc, "direct", true, false, false, false, nil)

	if err != nil {
		return err
	}

	// Declare the dead-letter queue:
	deadLetterQueue, err := ch.QueueDeclare(deadQ, true, false, false, false, nil)

	if err != nil {
		return err
	}

	// Bind the dead-letter queue to the dead_letter exchange with a binding key:
	err = ch.QueueBind(deadLetterQueue.Name, deadQKey, deadExc, false, nil)
	if err != nil {
		return err
	}

	return nil
}

func checkRecipient(recipientType string, value string) error {
	valid := true

	switch recipientType {
	case "email":
		// Use a regular expression to check if the email address is valid
		regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		valid = regex.MatchString(value)
	case "phone":
		// Use a regular expression to check if the phone number is valid
		regex := regexp.MustCompile(`^[0-9]{10,}$`)
		valid = regex.MatchString(value)
	default:
		// Unknown recipient type
		return fmt.Errorf("unknown recipient type: %s", recipientType)
	}

	if !valid {
		return fmt.Errorf("invalid %s: %s", recipientType, value)
	}

	return nil
}
