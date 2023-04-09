package main

import (
	"errors"
	"fmt"
)

type Action interface {
	Execute() error
	Rollback() error
	GetName() string
}

type Saga struct {
	Actions []Action
}

func (s *Saga) AddAction(a Action) {
	s.Actions = append(s.Actions, a)
}

func (s *Saga) Run() error {
	for i, action := range s.Actions {
		err := action.Execute()

		if err != nil {
			fmt.Println("Failed to execute " + action.GetName())

			for j := i - 1; j >= 0; j-- {
				err2 := s.Actions[j].Rollback()

				if err2 != nil {
					fmt.Println("TODO: Find way to handle this")
				}
			}

			return err
		}

		fmt.Println("Succeeded " + action.GetName())
		fmt.Println("---------")
	}

	return nil
}

// Action1
type Action1 struct {
	Name string
}

func (a *Action1) Execute() error {
	fmt.Println("Executing action " + a.Name)
	return nil
}

func (a *Action1) Rollback() error {
	fmt.Println("Rollback action " + a.Name)
	return nil
}

func (a *Action1) GetName() string {
	return a.Name
}

// Action2
type Action2 struct {
	Name string
}

func (a *Action2) Execute() error {
	fmt.Println("Executing action " + a.Name)
	return errors.New("unexpected error")
}

func (a *Action2) Rollback() error {
	fmt.Println("Rollback action " + a.Name)
	return nil
}

func (a *Action2) GetName() string {
	return a.Name
}

// Action3
type Action3 struct {
	Name string
}

func (a *Action3) Execute() error {
	fmt.Println("Executing action " + a.Name)
	return nil
}

func (a *Action3) Rollback() error {
	fmt.Println("Rollback action " + a.Name)
	return nil
}

func (a *Action3) GetName() string {
	return a.Name
}

func main() {
	saga := Saga{}

	action1 := &Action1{"Action1"}
	action2 := &Action2{"Action2"}
	action3 := &Action3{"Action3"}

	saga.AddAction(action1)
	saga.AddAction(action2)
	saga.AddAction(action3)

	saga.Run()
}
