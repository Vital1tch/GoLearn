package types

import "fmt"

type Truck struct {
	load        int64
	fuel        float64
	destination string
}

func (t Truck) Deliver() {
	fmt.Printf("Грузовик загружен на %d. Отправляется в %s. Оставшиеся топливо: %.1f", t.load, t.destination, t.fuel)
}

func (t Truck) Speed() int {
	return 200
}

func NewTruck(load int64, fuel float64, destination string) *Truck {
	return &Truck{
		load:        load,
		fuel:        fuel,
		destination: destination,
	}
}
