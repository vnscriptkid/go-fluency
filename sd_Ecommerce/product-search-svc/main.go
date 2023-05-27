package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
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

var es *elasticsearch.Client

func main() {
	go runApi()

	// Initialize Kafka consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:29092"}, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	topic := "product-created"

	// Subscribe to the 'product-created' topic
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize Elasticsearch client
	_es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	if err != nil {
		log.Fatal(err)
	}

	es = _es

	// Process messages
	for msg := range partitionConsumer.Messages() {
		var product Product
		err = json.Unmarshal(msg.Value, &product)
		if err != nil {
			log.Printf("Failed to unmarshal product event: %v\n", err)
			continue
		}

		// Parse and potentially enrich message here
		req := esapi.IndexRequest{
			Index:      "products",
			DocumentID: product.ID,
			Body:       strings.NewReader(string(msg.Value)),
			Refresh:    "true", // Set to "true" to make the document searchable immediately
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error indexing document: %s", err)
			continue
		}

		if res.IsError() {
			log.Fatalf("Elasticsearch error response: %s", res.Status())
		} else {
			log.Println("Document indexed successfully")
		}

		res.Body.Close()
	}
}

func runApi() {
	r := gin.Default()

	// GET /products/:id
	r.GET("/products/:id", getProduct)

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

func getProduct(c *gin.Context) {
	productID := c.Param("id")

	// Retrieve the product from Elasticsearch
	product, err := getProductFromElasticsearch(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func getProductFromElasticsearch(productID string) (Product, error) {
	req := esapi.GetRequest{
		Index:      "products",
		DocumentID: productID,
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		return Product{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return Product{}, err
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return Product{}, err
	}

	// Extract the product from the response data
	product := Product{
		ID:          data["_id"].(string),
		Name:        data["_source"].(map[string]interface{})["name"].(string),
		Description: data["_source"].(map[string]interface{})["description"].(string),
		Category:    data["_source"].(map[string]interface{})["category"].(string),
		Price:       data["_source"].(map[string]interface{})["price"].(float64),
	}

	return product, nil
}
