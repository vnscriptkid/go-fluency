package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

func main() {
	brokerList := []string{"localhost:29092"} // Replace with your Kafka broker address
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	client, err := sarama.NewClient(brokerList, config)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Error creating consumer: %s", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing consumer: %s", err)
		}
	}()

	topic := "pipe-and-filter"
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error consuming partition: %s", err)
	}

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			text := string(msg.Value)
			if strings.Contains(text, "3") {
				fmt.Println("Filtered message:", text)
			}
		case err := <-partitionConsumer.Errors():
			log.Fatalf("Error consuming messages: %s", err)
		}
	}
}
