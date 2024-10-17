package vehicle

type Bike struct {
	MaxSpeed   int
	TypeofBike string
}

func (b Bike) MaxSpeedOfVehicle() int {
	return b.MaxSpeed
}

func (b Bike) Type() string {
	return b.TypeofBike
}
