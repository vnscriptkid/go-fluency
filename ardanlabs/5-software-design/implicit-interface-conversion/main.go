package main

import "fmt"

type Mover interface {
	Move()
}
type Locker interface {
	Lock()
	Unlock()
}
type MoveLocker interface {
	Mover
	Locker
}

// Concrete type bike that implements all 3 interfaces
type bike struct{}

func (bike) Move() {
	fmt.Println("Moving the bike")
}
func (bike) Lock() {
	fmt.Println("Locking the bike")
}
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var moveLocker MoveLocker
	var mover Mover

	moveLocker = bike{}
	mover = moveLocker

	// cannot use mover (variable of type Mover) as MoveLocker value in assignment: Mover does not implement MoveLocker (missing method Lock)compilerInvalidIfaceAssign
	// moveLocker = mover

	fmt.Printf("%#v\n", moveLocker)
	fmt.Printf("%#v\n", mover)

	mover.Move()

	moveLocker.Lock()
	moveLocker.Unlock()
	moveLocker.Move()
}
