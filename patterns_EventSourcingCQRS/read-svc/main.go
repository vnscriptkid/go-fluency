package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type TransactionModel struct {
	ID        uuid.UUID `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type Balance struct {
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

func calculateBalance(user string, db *sql.DB) (float64, error) {
	var balance float64

	row := db.QueryRowContext(
		context.Background(),
		"SELECT current_balance FROM balances WHERE user_id = $1", user)

	err := row.Scan(&balance)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func updateBalance(user string, balance float64, db *sql.DB) error {
	_, err := db.ExecContext(
		context.Background(),
		`INSERT INTO 
			balances (user_id, current_balance) 
		VALUES 
			($1, $2) 
		ON CONFLICT (user_id) DO UPDATE SET current_balance = $2`, user, balance)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/banking")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transactions",
		GroupID: "group",
	})

	for {
		message, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Unable to read message from Kafka: %v\n", err)
		}

		var transaction TransactionModel
		err = json.Unmarshal(message.Value, &transaction)
		if err != nil {
			log.Fatalf("Unable to unmarshal transaction from JSON: %v\n", err)
		}

		fromUserBalance, err := calculateBalance(transaction.From, db)
		if err != nil {
			log.Fatalf("Unable to calculate balance for from user: %v\n", err)
		}
		fromUserBalance -= transaction.Amount
		err = updateBalance(transaction.From, fromUserBalance, db)
		if err != nil {
			log.Fatalf("Unable to update balance for from user: %v\n", err)
		}

		toUserBalance, err := calculateBalance(transaction.To, db)
		if err != nil {
			log.Fatalf("Unable to calculate balance for to user: %v\n", err)
		}
		toUserBalance += transaction.Amount
		err = updateBalance(transaction.To, toUserBalance, db)
		if err != nil {
			log.Fatalf("Unable to update balance for to user: %v\n", err)
		}
	}
}
