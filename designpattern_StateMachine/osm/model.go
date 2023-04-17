package osm

type Order struct {
	Status OrderStatus
}

type OrderStatus int

const (
	New OrderStatus = iota
	Confirmed
	Shipped
	Delivered
)

func (o *Order) SetStatus(status OrderStatus) {
	o.Status = status
}

func (o *Order) GetCurrentStatus() OrderStatus {
	return o.Status
}

type OrderState interface {
	Confirm(*Order)
	Ship(*Order)
	Deliver(*Order)
}
