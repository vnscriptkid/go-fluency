package main

import (
	"encoding/json"
	"log"
	"runtime/debug"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	orderDLQTopic = "orders-dlq"
)

type Order struct {
	ID          string `json:"id"`
	ProductID   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	OrderStatus string `json:"order_status"`
}

type DLQMessage struct {
	OriginalMessage json.RawMessage `json:"original_message"`
	ErrorMessage    string          `json:"error_message"`
	ErrorStackTrace string          `json:"error_stack_trace"`
	Timestamp       int64           `json:"timestamp"`
	ServiceName     string          `json:"service_name"`
}

func sendToDLQ(producer *kafka.Producer, topic string, msgValue []byte, err error) {
	stackTrace := string(debug.Stack())

	dlqMessage := DLQMessage{
		OriginalMessage: msgValue,
		ErrorMessage:    "", // err.Error(),
		ErrorStackTrace: stackTrace,
		Timestamp:       time.Now().Unix(),
		ServiceName:     "order-processor",
	}

	dlqMessageBytes, err := json.Marshal(dlqMessage)
	if err != nil {
		log.Printf("Failed to marshal DLQ message: %s\n", err)
		return
	}

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          dlqMessageBytes,
	}, nil)
}

func main() {
	broker := "localhost:29092"
	groupID := "order-processor"

	// Create a Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
		// "socket.timeout.ms": 30000, // Increase the timeout to 30 seconds
		// "debug": "all", // Enable debug logging
	})

	if err != nil {
		log.Fatalf("Failed to create consumer: %s\n", err)
	}

	// Create a Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	// Subscribe to the orders topic
	err = consumer.Subscribe("orders", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to orders topic: %s\n", err)
	}

	for {
		msg, err := consumer.ReadMessage(1000 * time.Millisecond)
		if err == nil {
			order := Order{}
			err := json.Unmarshal(msg.Value, &order)
			if err != nil {
				log.Printf("Error unmarshaling message: %s\n", err)
				sendToDLQ(producer, orderDLQTopic, msg.Value, err)
				continue
			}

			// Simulate a failure for orders with odd quantities
			if order.Quantity%2 != 0 {
				log.Printf("Failed to process order with ID %s\n", order.ID)
				sendToDLQ(producer, orderDLQTopic, msg.Value, err)
			} else {
				log.Printf("Processed order with ID %s\n", order.ID)
			}
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
