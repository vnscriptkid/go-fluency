package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/go-redis/redis"
)

const (
	DbURL = "postgres://postgres:123456@localhost:5432/messages?sslmode=disable"
)

var redisClient *redis.Client
var pgClient *sql.DB

func connectDBs() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to redis")

	pgClient, err = sql.Open("postgres", DbURL)

	if err != nil {
		panic(err)
	}

	err = pgClient.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to pg")
}

func createDb() {
	_, err := pgClient.Query(`
	DROP TABLE mytable;
	
	CREATE TABLE IF NOT EXISTS mytable (
		id serial primary key,
		value varchar(255)
		);
		`)
	if err != nil {
		log.Fatalf("failed to create table %v", err)
	}
}

func GetFromPostgres(key string) (int, string) {
	keyInt, err := strconv.Atoi(key)

	if err != nil {
		panic(err)
	}

	rows, err := pgClient.Query(`
		SELECT * FROM mytable WHERE id = $1
	`, keyInt)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var value string

		err := rows.Scan(&id, &value)

		if err != nil {
			panic(err)
		}

		log.Printf("From DB: %v, %v", id, value)

		return keyInt, value
	}

	return 0, ""
}

func WriteToRedis(key string, value string) {
	err := redisClient.Set(key, value, 0).Err()

	if err != nil {
		log.Fatalf("failed setting key %v", err)
	}
}

func WriteToPostgres(key string, value string) {
	intKey, err := strconv.Atoi(key)

	if err != nil {
		log.Fatalf("failed to convert str to int %v", err)
	}

	result, err := pgClient.Exec(`
			INSERT INTO mytable (id, value) VALUES ($1, $2);
		`, intKey, value)

	if err != nil {
		log.Printf("err inserting to table %v", err)
		return
	}

	if id, err := result.LastInsertId(); err != nil {
		log.Printf("Persisted %v", id)
	}
}

func ClearCacheByKey(key string) {
	err := redisClient.Del(key).Err()

	if err != nil {
		fmt.Printf("failed to clear cache by key %v", err)
	}
}

func WriteAround(key string, value string) {
	WriteToPostgres(key, value)

	// delete from cache
	ClearCacheByKey(key)
}

func ReadInCaseWriteAround(key string) {
	value, err := redisClient.Get(key).Result()

	if err == nil && value != "" {
		log.Printf("[v] Cache HIT %v", value)
		return
	}

	log.Printf("[x] Cache MISS %v", value)

	_, value = GetFromPostgres(key)

	redisClient.Set(key, value, time.Hour)
}

func WriteBack(key string, value string) {
	// client 	cache	DB
	// write =====>|
	// 			   =====>|
	// read  ----->|
	//            x
	// read  ---->x|
	WriteToRedis(key, value)

	// main thread won't wait here but will return to client right away
	go WriteToPostgres(key, value)
}

func WriteThrough(key string, value string) {
	WriteToRedis(key, value)

	WriteToPostgres(key, value)
}

// func WriteThrough(key string, value string) {}
// func WriteAround(key string, value string) {}

func GetValueOfKey(key string) {
	value, err := redisClient.Get(key).Result()

	if err == nil && value != "" {
		log.Printf("[v] Cache HIT %v", value)
		return
	}

	if err == redis.Nil {
		log.Printf("[x] Cache MISS")
	} else if err != nil {
		log.Fatalf("err getting key %v", err)
	}

	keyInt, err := strconv.Atoi(key)

	if err != nil {
		panic(err)
	}

	rows, err := pgClient.Query(`
		SELECT * FROM mytable WHERE id = $1
	`, keyInt)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var value string

		err := rows.Scan(&id, &value)

		if err != nil {
			panic(err)
		}

		log.Printf("From DB: %v, %v", id, value)
	}
}

func ClearCache() {
	err := redisClient.FlushAll().Err()

	if err != nil {
		panic(err)
	}

	log.Println("Cache cleared")
}

func main() {
	connectDBs()
	// createDb()

	// WriteBack("1", "one")
	GetValueOfKey("1")
	// ClearCache()
}
