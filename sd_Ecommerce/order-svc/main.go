package main

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

var producer sarama.SyncProducer

// Order model
type Order struct {
	ID          string
	ItemID      string
	Quantity    int
	OrderStatus string
}

// Event model
type Event struct {
	ID        string
	EventType string
	Data      Order
}

func main() {
	// Connect to kafka
	brokerList := []string{"localhost:29092"} // Replace with your Kafka broker address

	_producer, err := sarama.NewSyncProducer(brokerList, nil)

	if err != nil {
		log.Fatalf("Error creating producer: %s", err)
	}

	producer = _producer

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing producer: %s", err)
		}
	}()

	r := gin.Default(":8080")

	r.POST("/orders", func(c *gin.Context) {
		var order Order
		c.BindJSON(&order)

		// Persist order to database here...

		// Then publish order_created event
		event := Event{
			ID:        order.ID,
			EventType: "order_created",
			Data:      order,
		}
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          event,
		}, nil)

		// Call inventory service and update inventory...

		// Then publish inventory_updated event
		event.EventType = "inventory_updated"
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          event,
		}, nil)
		c.JSON(200, order)
	})

	r.Run()
}
