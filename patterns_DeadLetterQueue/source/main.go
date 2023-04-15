package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Order struct {
	ID          string `json:"id"`
	ProductID   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	OrderStatus string `json:"order_status"`
}

func handleEvents(producer *kafka.Producer) {
	go func() {
		for event := range producer.Events() {
			switch e := event.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					log.Printf("Failed to deliver message: %v\n", e.TopicPartition.Error)
				}
			case kafka.Error:
				log.Printf("Producer error: %v\n", e)
			default:
				// Ignore other event types
			}
		}
	}()
}

func main() {
	broker := "localhost:29092"
	topic := "orders"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create producer: %s\n", err)
	}

	defer producer.Close()

	go handleEvents(producer)

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 100; i++ {
		time.Sleep(1 * time.Second)

		order := Order{
			ID:          fmt.Sprintf("order-%d", i),
			ProductID:   fmt.Sprintf("product-%d", rand.Intn(100)+1),
			Quantity:    rand.Intn(10) + 1,
			OrderStatus: "created",
		}

		orderBytes, err := json.Marshal(order)
		if err != nil {
			log.Printf("Failed to marshal order: %s\n", err)
			continue
		}

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          orderBytes,
		}, nil)

		if err != nil {
			continue
		}

		log.Printf("Sent order: %s\n", order.ID)
	}

	producer.Flush(15 * 1000)
}
