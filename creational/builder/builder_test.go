package builder

import "testing"

func TestBuilder(t *testing.T){
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	if car.Structure != "car"{
		t.Errorf("Expected car type, instead got %s", car.Structure)
	}
	if car.Wheels != 4{
		t.Errorf("Expected 4 wheels, instead got %d", car.Wheels)
	}
	if car.Seats != 5{
		t.Errorf("Expected 5 seats, instead got %d", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorBike := bikeBuilder.GetVehicle()
	if motorBike.Structure != "motorbike"{
		t.Errorf("Expected motorBike type, instead got %s", motorBike.Structure)
	}
	if motorBike.Wheels != 2{
		t.Errorf("Expected 4 wheels, instead got %d", motorBike.Wheels)
	}
	if motorBike.Seats != 2{
		t.Errorf("Expected 5 seats, instead got %d", motorBike.Seats)
	}

}