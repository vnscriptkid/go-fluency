package osm

import "fmt"

type ShippedOrder struct{}

func (s *ShippedOrder) Confirm(o *Order) {
	fmt.Println("Order has already been confirmed")
}

func (s *ShippedOrder) Ship(o *Order) {
	fmt.Println("Order has already been shipped")
}

func (s *ShippedOrder) Deliver(o *Order) {
	o.SetStatus(Delivered)
}
