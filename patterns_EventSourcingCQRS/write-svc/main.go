package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

func addTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	transaction.ID = uuid.New()
	transaction.Timestamp = time.Now()

	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/banking")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	_, err = db.ExecContext(context.Background(), `
		INSERT INTO transactions 
			(id, from_user_id, to_user_id, amount, timestamp) 
		VALUES ($1, $2, $3, $4, $5)`,
		transaction.ID, transaction.From, transaction.To, transaction.Amount, transaction.Timestamp)

	if err != nil {
		log.Fatalf("Unable to insert transaction: %v\n", err)
	}

	// Publish transaction to Kafka
	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
	})

	message := kafka.Message{
		Key:   []byte(transaction.From),
		Value: toJSON(transaction),
	}

	err = kafkaWriter.WriteMessages(context.Background(), message)

	if err != nil {
		log.Fatalf("Unable to publish message to Kafka: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/transactions", addTransactionHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
