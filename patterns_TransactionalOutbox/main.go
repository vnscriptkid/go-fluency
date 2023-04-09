package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Message struct {
	ID        int       `json:"id"`
	Payload   string    `json:"payload"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	SentAt    time.Time `json:"sent_at"`
}

func send(ctx context.Context, id int, payload string) error {
	fmt.Printf("Sending to KAFKA msg ID [%v] with payload [%v]\n", id, payload)
	return nil
}

type dbClient struct {
	client *sql.DB
}

func (d *dbClient) InitDb() {
	_, err := d.client.Exec(`
		CREATE TABLE IF NOT EXISTS outbox (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			payload JSON NOT NULL,
			status TEXT NOT NULL DEFAULT 'unsent',
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			sent_at TIMESTAMP DEFAULT NULL
		);
		
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)

	if err != nil {
		log.Fatal("failed to create table")
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./db/outbox.db")

	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()

	dbC := dbClient{client: db}

	dbC.InitDb()

	ctx := context.Background()
	msgCh := make(chan Message)
	errCh := make(chan error)

	go dbC.SimulateDbInserts(context.Background())

	go func() {
		// Infinitely listen from channel for msg from `msgCh`
		// and err from `errCh`
		for {
			select {
			case msg := <-msgCh:
				err := send(ctx, msg.ID, msg.Payload)

				// Failure can happen here
				// After sending succeeded, but it fails before updating outbox
				// So send() here guarantees At least once
				// It may need idempotency

				if err != nil {
					errCh <- err
				} else {
					_, err := db.Exec("UPDATE outbox SET status = ?, sent_at = CURRENT_TIMESTAMP WHERE id = ?", "sent", msg.ID)
					if err != nil {
						errCh <- err
					}
				}
			case err := <-errCh:
				log.Println(err)
				time.Sleep(time.Second)
			}
		}
	}()

	// Constantly polling from `outbox` table for unsent messages
	for {
		// fmt.Println("Polling from `outbox`")
		tx, err := db.BeginTx(ctx, nil)

		if err != nil {
			log.Printf("failed to begin tx %v", err)
			time.Sleep(time.Second)
			continue
		}

		rows, err := tx.QueryContext(ctx, "SELECT id, payload FROM outbox WHERE status = ?", "unsent")

		if err != nil {
			log.Printf("failed to query from outbox %v", err)
			time.Sleep(time.Second)
			continue
		}

		for rows.Next() {
			var msg Message
			if err := rows.Scan(&msg.ID, &msg.Payload); err != nil {
				log.Printf("failed to scan data %v", err)
				continue
			}

			// waits for either a message to be sent to `msgCh`` or for 10 seconds to elapse,
			// whichever happens first.
			select {
			case msgCh <- msg:
			case <-time.After(10 * time.Second):
				log.Printf("Timed out sending message: %+v\n", msg)
			}
		}

		if err := rows.Close(); err != nil {
			log.Printf("failed to close on rows")
			time.Sleep(time.Second)
			continue
		}

		if err := tx.Commit(); err != nil {
			log.Printf("failed to commit tx")
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(time.Second)
	}
}

func (d *dbClient) SimulateDbInserts(ctx context.Context) {
	i := 1

	for {
		time.Sleep(time.Second)
		tx, err := d.client.BeginTx(ctx, nil)

		if err != nil {
			fmt.Printf("failed to begin tx %v", err)
			continue
		}

		randomName := "user-" + fmt.Sprint(i)
		i++

		result, err := tx.Exec("INSERT INTO users (name) VALUES (?)", randomName)

		if err != nil {
			fmt.Printf("failed to insert to users: %v", err)
			err := tx.Rollback()

			if err != nil {
				fmt.Printf("failed to rollback")
			}
			continue
		}

		effected, _ := result.RowsAffected()
		fmt.Printf("inserted to users: %v\n", effected)

		result, err = tx.Exec("INSERT INTO outbox (payload) VALUES (?)", randomName)

		if err != nil {
			fmt.Printf("failed to insert to outbox: %v", err)
			err := tx.Rollback()

			if err != nil {
				fmt.Printf("failed to rollback")
			}
			continue
		}

		effected, _ = result.RowsAffected()
		fmt.Printf("inserted to outbox: %v\n", effected)

		err = tx.Commit()

		if err != nil {
			fmt.Printf("failed to commit %v", err)
		}
	}
}
