package main

import (
	"fmt"
	"log"

	common "Patterns_SagaChoreography/common"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker        = "localhost:9092"
	paymentTopic  = "payment"
	shippingTopic = "shipping"
	statusTopic   = "status"
)

func chargePaymentHandler(producer *kafka.Producer, message string) {
	log.Printf("Charge Payment: %s", message)

	// Send message to the next service
	common.SendMessage(producer, shippingTopic, fmt.Sprintf("%s: Ship Order", message))

	// Send status update
	common.SendMessage(producer, statusTopic, fmt.Sprintf("%s: Inventory Reserved", message))
}

func main() {
	paymentConsumer := common.NewKafkaConsumer(broker, "payment_group", paymentTopic)

	common.ConsumeEvents(paymentConsumer, func(s string) {
		chargePaymentHandler(common.NewKafkaProducer(broker), s)
	})
}
