package abstract_factory

type Sportmotorbike struct{}

func (*Sportmotorbike) NumWheels() int {
	return 2
}
func (*Sportmotorbike) NumSeats() int {
	return 1
}
func (*Sportmotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
