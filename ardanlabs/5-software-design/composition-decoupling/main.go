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

// CustomOrderProcessor handles processing orders
type CustomOrderProcessor struct {
	OrdersProcessed int
}

// ProcessOrder processes an order
func (op *CustomOrderProcessor) ProcessOrder(o *Order) error {
	fmt.Printf("Processing order ID: %s\n", o.ID)
	// Simulate checking inventory and processing the order
	time.Sleep(1 * time.Second) // Simulate delay
	o.OrderStatus = "Processed"
	op.OrdersProcessed++
	return nil
}

type OrderProcessor interface {
	ProcessOrder(o *Order) error
}

func ProcessOrder(orderProcessor OrderProcessor, o *Order) error {
	return orderProcessor.ProcessOrder(o)
}

// CustomShipper handles shipping orders
type CustomShipper struct {
	OrdersShipped int
}

// ShipOrder ships an order
func (s *CustomShipper) ShipOrder(o *Order) error {
	fmt.Printf("Shipping order ID: %s\n", o.ID)
	// Simulate shipping process
	time.Sleep(1 * time.Second) // Simulate delay
	o.OrderStatus = "Shipped"
	s.OrdersShipped++
	return nil
}

type Shipper interface {
	ShipOrder(o *Order) error
}

func ShipOrder(shipper Shipper, o *Order) error {
	return shipper.ShipOrder(o)
}

// LogisticsSystem `composes` CustomOrderProcessor and CustomShipper to manage the logistics
// This struct combines the functionalities of both CustomOrderProcessor and CustomShipper by embedding these types
// It represents a higher-level system that can manage the complete lifecycle of an order from processing to shipping.
type LogisticsSystem struct {
	CustomOrderProcessor
	CustomShipper
}

// Composite interface that combines OrderProcessor and Shipper
type OrderProcessorShipper interface {
	OrderProcessor
	Shipper
}

// NewLogisticsSystem creates a new instance of LogisticsSystem
func NewLogisticsSystem() *LogisticsSystem {
	return &LogisticsSystem{}
}

// ManageOrder takes an order and handles its processing and shipping
func ManageOrder(orderProcessorShipper OrderProcessorShipper, o *Order) error {
	if err := ProcessOrder(orderProcessorShipper, o); err != nil {
		return err
	}
	if err := ShipOrder(orderProcessorShipper, o); err != nil {
		return err
	}
	return nil
}

func main() {
	ls := NewLogisticsSystem()
	order := &Order{ID: "1234", Item: "Book", Quantity: 1}

	ManageOrder(ls, order)
}
