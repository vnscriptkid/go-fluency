package osm

import "fmt"

type DeliveredOrder struct{}

func (s *DeliveredOrder) Confirm(o *Order) {
	fmt.Println("Order has already been confirmed")
}

func (s *DeliveredOrder) Ship(o *Order) {
	fmt.Println("Order has already been shipped")
}

func (s *DeliveredOrder) Deliver(o *Order) {
	fmt.Println("Order has already been delivered")
}
