package types

import "fmt"

type Ship struct {
	capacity        int64
	currentLocation string
	destination     string
}

func (s Ship) Deliver() {
	fmt.Printf("Корабль имеет вместимость: %d, находится в %s, а отправляется в %s\n", s.capacity, s.currentLocation, s.destination)
}

func (s Ship) Speed() int {
	return 320
}

func NewShip(capacity int64, currentLocation, destination string) *Ship {
	return &Ship{
		capacity:        capacity,
		currentLocation: currentLocation,
		destination:     destination,
	}
}
