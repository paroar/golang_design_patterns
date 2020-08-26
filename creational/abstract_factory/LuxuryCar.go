package abstract_factory

type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
	return 3
}
func (*LuxuryCar) NumWheels() int {
	return 4
}
func (*LuxuryCar) NumSeats() int {
	return 2
}