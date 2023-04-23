package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var dbShards []*sql.DB

func initDB() {
	// Connection string for the shards
	connStrs := []string{
		"postgres://user:password@localhost:5433/shard?sslmode=disable",
		"postgres://user:password@localhost:5434/shard?sslmode=disable",
	}

	for _, connStr := range connStrs {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		dbShards = append(dbShards, db)
	}
}

func shardHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Choose shard based on user_id
	shardIndex := userId % len(dbShards)
	db := dbShards[shardIndex]

	// Perform database operation, for example, insert user data
	_, err = db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", userId, "John Doe")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data saved to shard %d", shardIndex)
}

func main() {
	initDB()

	http.HandleFunc("/", shardHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
