// client.go
package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		if messageType == websocket.TextMessage {
			fmt.Printf("received: %s\n", message)
		} else if messageType == websocket.CloseMessage {
			break
		}
	}
}
