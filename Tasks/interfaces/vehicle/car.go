package vehicle

type Car struct {
	MaxSpeed  int
	TypeofCar string
}

func (c Car) MaxSpeedOfVehicle() int {
	return c.MaxSpeed
}

func (c Car) Type() string {
	return c.TypeofCar
}
