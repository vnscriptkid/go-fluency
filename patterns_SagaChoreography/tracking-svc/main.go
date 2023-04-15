package main

import (
	common "Patterns_SagaChoreography/common"
	"log"
	"strings"
	"sync"
)

const (
	broker      = "localhost:9092"
	statusTopic = "status"
)

type OrderStatus struct {
	sync.RWMutex
	statusMap map[string]string
}

func (os *OrderStatus) updateStatus(orderID, status string) {
	os.Lock()
	defer os.Unlock()
	os.statusMap[orderID] = status
}

func (os *OrderStatus) getStatus(orderID string) string {
	os.RLock()
	defer os.RUnlock()
	return os.statusMap[orderID]
}

func statusHandler(orderStatus *OrderStatus, message string) {
	log.Printf("Received status update: %s", message)

	// Parse the message to extract order ID and status
	orderID, status := parseMessage(message)

	// Update the status in the OrderStatus instance
	orderStatus.updateStatus(orderID, status)

	// In a real-world scenario, you might expose an API for clients to query the status of a process
}

func parseMessage(message string) (string, string) {
	// Implement a function to parse the message and extract the order ID and status
	// Assuming the message format is "OrderID: Status"
	// You can adjust the implementation based on the actual message format
	parts := strings.SplitN(message, ": ", 2)

	if len(parts) != 2 {
		return "", ""
	}

	return parts[0], parts[1]
}

func main() {
	statusConsumer := common.NewKafkaConsumer(broker, "status_group", statusTopic)
	orderStatus := &OrderStatus{statusMap: make(map[string]string)}

	common.ConsumeEvents(statusConsumer, func(message string) { statusHandler(orderStatus, message) })
}
