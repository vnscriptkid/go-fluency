package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	clients   = make(map[*websocket.Conn]bool)
	clientMtx = sync.Mutex{}
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	go broadcastTime()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	clientMtx.Lock()
	clients[conn] = true
	clientMtx.Unlock()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			clientMtx.Lock()
			delete(clients, conn)
			clientMtx.Unlock()
			break
		}
	}
}

func broadcastTime() {
	for {
		time.Sleep(1 * time.Second)

		message := fmt.Sprintf("Current server time: %s", time.Now().Format(time.RFC3339))
		clientMtx.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("write error:", err)
				delete(clients, conn)
			}
		}
		clientMtx.Unlock()
	}
}
