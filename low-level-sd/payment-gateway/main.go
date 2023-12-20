package main

import "fmt"

type StateValue int

const (
	StateValueSuccessful StateValue = iota
	StateValueFailed
	StateValueInProgress
	StateValueOnHold
)

type State interface {
	ChangeFrom(prev StateValue)
}

type StateSuccessful struct {
	Value StateValue
}

// func (s *StateSuccessful) ChangeTo(next StateValue) {
// 	switch next {
// 	case StateValueFailed:
// 		fmt.Println("From %v to %v", s.Value, next)
// 	case StateValueOnHold:
// 		fmt.Println("Unexpected rollback from %v to %v", s.Value, next)
// 	}
// }

func (s *StateSuccessful) ChangeFrom(from StateValue) {
	switch from {
	case StateValueFailed:
		fmt.Println("From %v to %v", from, s.Value)
	case StateValueOnHold:
		fmt.Println("Unexpected rollback from %v to %v", from, s.Value)
	}
}

func NewStateSuccessful() *StateSuccessful {
	return &StateSuccessful{
		Value: StateValueSuccessful,
	}
}

func (s *Transaction) setState(newState StateValue) {
	s.Status = newState

}

type User struct {
	ID int
}

type Bank struct {
	ID      int
	Name    string
	Address string
}

type Account struct {
	ID   int
	Bank Bank
	User User
}

type Transaction struct {
	ID       int
	Status   StateValue
	Sender   Account
	Receiver Account
	Amount   float64
	Currency string
}

func main() {
	// t := Transaction{}
	// t.SetState(StateInProgress)

	// fmt.Println(t)

	t := Transaction{
		Status: NewStateSuccessful(),
	}

}
