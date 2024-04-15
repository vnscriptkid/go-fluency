package main

import (
	"fmt"
	"time"
)

// Data structure for holding order details
type Order struct {
	ID          string
	Item        string
	Quantity    int
	OrderStatus string
}

// OrderProcessor handles processing orders
type OrderProcessor struct {
	OrdersProcessed int
}

// ProcessOrder processes an order
func ProcessOrder(orderProcessor *OrderProcessor, o *Order) error {
	fmt.Printf("Processing order ID: %s\n", o.ID)
	// Simulate checking inventory and processing the order
	time.Sleep(1 * time.Second) // Simulate delay
	o.OrderStatus = "Processed"
	orderProcessor.OrdersProcessed++
	return nil
}

// Shipper handles shipping orders
type Shipper struct {
	OrdersShipped int
}

// ShipOrder ships an order
func ShipOrder(shipper *Shipper, o *Order) error {
	fmt.Printf("Shipping order ID: %s\n", o.ID)
	// Simulate shipping process
	time.Sleep(1 * time.Second) // Simulate delay
	o.OrderStatus = "Shipped"
	shipper.OrdersShipped++
	return nil
}

// LogisticsSystem composes OrderProcessor and Shipper to manage the logistics
// This struct combines the functionalities of both OrderProcessor and Shipper by embedding these types
// It represents a higher-level system that can manage the complete lifecycle of an order from processing to shipping.
type LogisticsSystem struct {
	OrderProcessor
	Shipper
}

// NewLogisticsSystem creates a new instance of LogisticsSystem
func NewLogisticsSystem() *LogisticsSystem {
	return &LogisticsSystem{}
}

func ManageOrder(ls *LogisticsSystem, o *Order) error {
	if err := ProcessOrder(&ls.OrderProcessor, o); err != nil {
		return err
	}
	if err := ShipOrder(&ls.Shipper, o); err != nil {
		return err
	}
	return nil
}

func main() {
	ls := NewLogisticsSystem()
	order := &Order{ID: "1234", Item: "Book", Quantity: 1}

	ManageOrder(ls, order)
}
