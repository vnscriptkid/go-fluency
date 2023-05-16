package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

func main() {
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

	// Subscribe to the 'client-data' topic
	partitionConsumer, err := consumer.ConsumePartition("client-data", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize Elasticsearch client
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(es.Info)

	// Process messages
	for msg := range partitionConsumer.Messages() {
		// Parse and potentially enrich message here
		req := esapi.IndexRequest{
			Index:   "client-data",
			Body:    strings.NewReader(string(msg.Value)),
			Refresh: "true", // Set to "true" to make the document searchable immediately
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
