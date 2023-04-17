package osm

import "fmt"

type ConfirmedOrder struct{}

func (s *ConfirmedOrder) Confirm(o *Order) {
	fmt.Println("Order has already been confirmed")
}

func (s *ConfirmedOrder) Ship(o *Order) {
	o.SetStatus(Shipped)
}

func (s *ConfirmedOrder) Deliver(o *Order) {
	fmt.Println("Cannot deliver order that has not been shipped")
}
