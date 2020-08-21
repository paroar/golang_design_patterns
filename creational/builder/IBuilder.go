package builder

type Builder interface {
	SetProduct() Builder
	SetWheels() Builder
	SetSeats() Builder
	GetVehicle() VehicleProduct
}
