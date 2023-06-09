package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
)

var producer sarama.SyncProducer

type LocationUpdateRequest struct {
	DriverID int     `json:"driver_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
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

	r := gin.Default()
	r.POST("/location", func(c *gin.Context) {
		// Parse the driver's ID and new location from the request body
		var request LocationUpdateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Publish the new location to Kafka
		err = publishDriverLocation(&request)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "location updated"})
	})
	r.Run(":8080")
}

func publishDriverLocation(location *LocationUpdateRequest) error {
	topic := "driver-locations"

	// Convert client data to JSON for sending to Kafka
	clientDataJSON, err := json.Marshal(location)

	if err != nil {
		log.Printf("Failed to marshal product to JSON: %v\n", err)
		return err
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(location.DriverID),
		Value: sarama.StringEncoder(clientDataJSON),
	})

	if err != nil {
		log.Printf("Error sending message: %s\n", err)
		return err
	}

	log.Printf("Published to %s: %s", topic, clientDataJSON)

	return nil
}
