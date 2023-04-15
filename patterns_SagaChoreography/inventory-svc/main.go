package main

import (
	"fmt"
	"log"

	common "Patterns_SagaChoreography/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker         = "localhost:9092"
	inventoryTopic = "inventory"
	paymentTopic   = "payment"
	statusTopic    = "status"
)

func reserveInventoryHandler(producer *kafka.Producer, message string) {
	log.Printf("Reserve Inventory: %s", message)

	// Send status update
	common.SendMessage(producer, statusTopic, fmt.Sprintf("%s: Inventory Reserved", message))

	// Send message to the next service
	common.SendMessage(producer, paymentTopic, fmt.Sprintf("%s: Charge Payment", message))
}

func main() {
	inventoryConsumer := common.NewKafkaConsumer(broker, "inventory_group", inventoryTopic)

	producer := common.NewKafkaProducer(broker)

	common.ConsumeEvents(inventoryConsumer, func(s string) {
		reserveInventoryHandler(producer, s)
	})
}
