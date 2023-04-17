package osm

import "fmt"

type NewOrder struct{}

func (s *NewOrder) Confirm(o *Order) {
	o.SetStatus(Confirmed)
}

func (s *NewOrder) Ship(o *Order) {
	fmt.Println("Cannot ship order that has not been confirmed")
}

func (s *NewOrder) Deliver(o *Order) {
	fmt.Println("Cannot deliver order that has not been shipped")
}
