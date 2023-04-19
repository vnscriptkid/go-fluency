package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	// Timeline
	// 0      1      2       3  ...
	// |   |  |      |
	// s1  p  s2     p

	// Create a Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6378",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to Redis")

	// Start a goroutine to handle incoming messages
	go func() {
		go func() {
			pubsub1 := rdb.Subscribe(context.Background(), "mychannel")
			defer pubsub1.Close()

			fmt.Println("[Sub-1] Subscribe to channel at timestamp 0")
			for {
				msg, err := pubsub1.ReceiveMessage(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("[Sub-1] Received message: %s\n", msg.Payload)
			}
		}()

		time.Sleep(2 * time.Second)

		go func() {
			pubsub2 := rdb.Subscribe(context.Background(), "mychannel")
			defer pubsub2.Close()

			fmt.Println("[Sub-2] Subscribe to channel at timestamp 2")
			for {
				msg, err := pubsub2.ReceiveMessage(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("[Sub-2] Received message: %s\n", msg.Payload)
			}
		}()
	}()

	// Publish a message to the channel
	go func() {
		time.Sleep(1 * time.Second)

		if err := rdb.Publish(context.Background(), "mychannel", "One").Err(); err != nil {
			log.Fatal("Failed to publish: ", err)
		} else {
			fmt.Println("Published message 1 at timestamp 1")
		}

		time.Sleep(2 * time.Second)

		if err := rdb.Publish(context.Background(), "mychannel", "Two").Err(); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Published message 2 at timestamp 3")
		}
	}()

	time.Sleep(5 * time.Second)
}
