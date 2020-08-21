package builder

type CarBuilder struct{
	v VehicleProduct
}

func (c *CarBuilder) SetProduct() Builder{
	c.v.Structure = "car"
	return c
}
func (c *CarBuilder) SetWheels() Builder{
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() Builder{
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct{
	return c.v
}
