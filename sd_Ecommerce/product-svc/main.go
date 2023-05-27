package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var producer sarama.SyncProducer
var productsCollection *mongo.Collection

func main() {
	// Create a MongoDB client by establishing a connection to the MongoDB server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the MongoDB server:
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	productsCollection = client.Database("products_db").Collection("products")

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

	r.POST("/products", createProduct)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}

func createProduct(c *gin.Context) {
	// Parse request body
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set created and updated timestamps
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	product.ID = uuid.New().String()

	// Save the product to MongoDB
	_, err := productsCollection.InsertOne(c, product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Publish product event
	err = publishProductEvent(product)
	if err != nil {
		log.Printf("Failed to publish product event: %v\n", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created", "ID": product.ID})
}

func publishProductEvent(product Product) error {
	topic := "product-created"

	// Convert client data to JSON for sending to Kafka
	clientDataJSON, err := json.Marshal(product)

	if err != nil {
		log.Printf("Failed to marshal product to JSON: %v\n", err)
		return err
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(clientDataJSON),
	})

	if err != nil {
		log.Printf("Error sending message: %s\n", err)
		return err
	}

	log.Printf("Published to %s: %s", topic, clientDataJSON)

	return nil
}
