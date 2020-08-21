package builder

type BikeBuilder struct{
	v VehicleProduct
}

func (b *BikeBuilder) SetProduct() Builder{
	b.v.Structure = "motorbike"
	return b
}
func (b *BikeBuilder) SetWheels() Builder{
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() Builder{
	b.v.Seats = 2
	return b
}
func (c *BikeBuilder) GetVehicle() VehicleProduct{
	return c.v
}
