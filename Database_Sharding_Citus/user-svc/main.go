package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	// Connection string for the Citus master node
	connStr := "postgres://user:password@localhost:5433/postgres?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func shardHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	// Perform database operation, for example, insert user data
	_, err = db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", userId, "John Doe")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data saved to shard")
}

func main() {
	initDB()

	http.HandleFunc("/", shardHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
