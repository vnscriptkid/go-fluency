package main

import (
	osm "designpattern_StateMachine/osm"
)

func main() {
	// Create a new order
	order := &osm.Order{
		Status: osm.New,
	}

	// Confirm the order
	orderState1 := &osm.NewOrder{}
	orderState1.Confirm(order)

	// Ship the order
	orderState2 := &osm.ConfirmedOrder{}
	orderState2.Ship(order)

	// Deliver the order
	orderState3 := &osm.ShippedOrder{}
	orderState3.Deliver(order)

	// Try to ship the order again
	orderState4 := &osm.ShippedOrder{}
	orderState4.Ship(order) // output: "Order has already been shipped"
}
