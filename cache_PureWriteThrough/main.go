package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

type Connection struct {
	RedisClient *redis.Client
	Db          *sql.DB
}

func (conn *Connection) close() {
	conn.RedisClient.Close()
	conn.Db.Close()
}

func (conn *Connection) updateConfigurationInDatabase(key string, value int) error {
	// Check if the configuration already exists in the database
	var count int
	err := conn.Db.QueryRow(`SELECT COUNT(*) FROM mytable WHERE name = $1`, key).Scan(&count)
	if err != nil {
		return err
	}

	// If the configuration does not exist, insert it
	if count == 0 {
		_, err = conn.Db.Exec(`INSERT INTO mytable (name, score) VALUES ($1, $2)`, key, value)
		if err != nil {
			return err
		}
	} else {
		// If the configuration exists, update it
		_, err = conn.Db.Exec(`UPDATE mytable SET score = $1 WHERE name = $2`, value, key)
		if err != nil {
			return err
		}
	}

	return nil
}

// This function updates the configuration in both the database and the cache.
func (conn *Connection) UpdateConfiguration(key string, value int) error {
	// Update the configuration in the database
	err := conn.updateConfigurationInDatabase(key, value)
	if err != nil {
		return err
	}

	// Update the configuration in the Redis cache
	cacheKey := fmt.Sprintf("configuration:%s", key)
	err = conn.RedisClient.Set(ctx, cacheKey, value, 0).Err()

	return err
}

// This function retrieves the configuration value directly from the cache.
// If the cache entry is not found (first time the configuration is requested),
// it retrieves the configuration from the database and stores it in the cache.
func (conn *Connection) GetConfiguration(key string) (string, error) {
	cacheKey := fmt.Sprintf("configuration:%s", key)
	value, err := conn.RedisClient.Get(ctx, cacheKey).Result()

	if err == redis.Nil {
		fmt.Println("!! Cache miss for key:", key)
		var score int

		r := conn.Db.QueryRow(`SELECT score FROM mytable WHERE name = $1`, key)

		if r.Err() != nil {
			return "", err
		}

		err = r.Scan(&score)

		if err != nil {
			return "", err
		}

		value = fmt.Sprintf("%d", score)

		err = conn.RedisClient.Set(ctx, cacheKey, value, 0).Err()

		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	fmt.Println("-- Cache hit for key:", key)

	return value, nil
}

func (conn *Connection) SimulateIntensiveReads() error {
	// Simulate intensive reads from the database
	for {
		for _, name := range []string{"james", "henry", "jerry"} {
			score, err := conn.GetConfiguration(name)

			if err != nil {
				fmt.Println("Error getting configuration:", err)
				continue
			}

			fmt.Println("Configuration:", name, score)
		}

		time.Sleep(1 * time.Second)
	}
}

func (conn *Connection) SimulateIntensiveWrites() error {
	// Simulate intensive writes to the database
	for {
		for _, name := range []string{"james", "henry", "jerry"} {
			// make sure randomized values are different from the previous ones
			rand.Seed(time.Now().UnixNano() + int64(name[0]))

			// random score from 10 to 20
			score := 10 + rand.Intn(10)

			conn.UpdateConfiguration(name, score)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
	})
	// Check if the Redis connection is working
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Could not connect to Redis:", err)
		return
	}

	fmt.Println("Redis connection successful:", pong)

	db, err := sql.Open("postgres", "postgres://postgres:123456@localhost:5440/mydb?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Database connection successful")

	conn := &Connection{
		RedisClient: redisClient,
		Db:          db,
	}
	defer conn.close()

	forever := make(chan bool)

	go conn.SimulateIntensiveReads()
	go conn.SimulateIntensiveWrites()

	forever <- true

}
