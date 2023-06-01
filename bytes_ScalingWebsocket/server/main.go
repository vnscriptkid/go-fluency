package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var redisClient *redis.Client
var ctx = context.Background()
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// check the origin of the request and return true if it's
		// from a trusted domain
		// if r.Header.Get("Origin") == "http://trusted.domain.com" {
		//     return true
		// }
		// return false
		return true
	},
}
var clients = make(map[string]*websocket.Conn)
var clientMtx = sync.Mutex{}

type User struct {
	ID string `json:"id"`
}

type Message struct {
	Sender    User   `json:"sender"`
	Recipient User   `json:"recipient"`
	Text      string `json:"text"`
}

func main() {
	var (
		REDIS_HOST = os.Getenv("REDIS_HOST")
		REDIS_PORT = os.Getenv("REDIS_PORT")
	)

	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT), // replace with your Redis server address
		Password: "",                                           // no password set
		DB:       0,                                            // use default DB
	})

	router := gin.Default()
	router.GET("/ws/:id", handleWebSocketConnection)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run("0.0.0.0:8080") // replace with your server address
}

func handleWebSocketConnection(c *gin.Context) {
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Failed to set websocket upgrade:", err)
		return
	}

	// Register server as subscriber to Redis channel `user.ID`
	user := User{ID: c.Param("id")}
	pubsub := redisClient.Subscribe(ctx, user.ID)

	// Save connection to clients map {userID:conn}
	clientMtx.Lock()
	clients[user.ID] = conn
	clientMtx.Unlock()

	go handleSocketMessage(conn, pubsub)
	go listenForPubsubMessages(conn, pubsub)
}

func handleSocketMessage(conn *websocket.Conn, pubsub *redis.PubSub) {
	defer conn.Close()
	for {
		_, messageBytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error on read message", err)
			return
		}

		var message Message
		err = json.Unmarshal(messageBytes, &message)
		if err != nil {
			log.Println("Error on unmarshal message", err)
			return
		}

		fmt.Printf("Received socket message: %+v\n", message)

		// Check if recipient is connected to server

		// If yes broadcast message to recipient via socket
		if recipientConn, ok := clients[message.Recipient.ID]; ok {
			err = recipientConn.WriteJSON(message)
			if err != nil {
				log.Println("Error on write message", err)
				continue
			}
			fmt.Printf("Recipient %s is connected to server\n", message.Recipient.ID)
			// Else likely recipient is connecting to another server, publish message to Redis pubsub channel `message.Recipient.ID`
		} else {
			err = redisClient.Publish(ctx, message.Recipient.ID, string(messageBytes)).Err()
			if err != nil {
				log.Println("Error on publish message", err)
				continue
			}
			fmt.Printf("Recipient %s is not connected to server\n", message.Recipient.ID)
		}
	}
}

func listenForPubsubMessages(conn *websocket.Conn, pubsub *redis.PubSub) {
	ch := pubsub.Channel()

	for msg := range ch {
		message := &Message{}
		err := json.Unmarshal([]byte(msg.Payload), message)
		if err != nil {
			log.Println("Error on unmarshal message", err)
			return
		}

		fmt.Printf("Received pubsub message: %+v\n", message)

		err = conn.WriteJSON(message)
		if err != nil {
			log.Println("Error on write message", err)
			return
		}
	}
}
