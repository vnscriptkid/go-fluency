package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	brokerList := []string{"localhost:29092"} // Replace with your Kafka broker address

	producer, err := sarama.NewSyncProducer(brokerList, nil)

	if err != nil {
		log.Fatalf("Error creating producer: %s", err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %s", err)
		}
	}()

	topic := "pipe-and-filter"

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Message %d", i)
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(msg),
		})
		if err != nil {
			log.Fatalf("Error sending message: %s", err)
		}
		time.Sleep(1 * time.Second)
	}
}
