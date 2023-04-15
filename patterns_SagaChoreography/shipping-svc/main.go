package main

import (
	"fmt"
	"log"

	common "Patterns_SagaChoreography/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker        = "localhost:9092"
	shippingTopic = "shipping"
	statusTopic   = "status"
)

func shipOrderHandler(producer *kafka.Producer, message string) {
	log.Printf("Ship Order: %s", message)
	// You can send a message to another topic if needed, e.g., to notify that the order has been shipped.

	// Send status update
	common.SendMessage(producer, statusTopic, fmt.Sprintf("%s: Inventory Reserved", message))

}

func main() {
	shippingConsumer := common.NewKafkaConsumer(broker, "shipping_group", shippingTopic)

	producer := common.NewKafkaProducer(broker)

	common.ConsumeEvents(shippingConsumer, func(message string) { shipOrderHandler(producer, message) })
}
