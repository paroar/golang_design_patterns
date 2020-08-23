package abstract_factory

import (
	"testing"
)

func TestBuildFactoryCar(t *testing.T) {
	carFactory, err := BuildFactory(CarFactoryType)

	if err != nil {
		t.Fatal("Factory not recognized")
	}

	//FamilyCar
	carVehicleF, fErr := carFactory.NewVehicle(FamilyCarType)
	if fErr != nil {
		t.Fatal(fErr)
	}
	familyCarSeats := carVehicleF.NumSeats()
	if familyCarSeats != 5 {
		t.Errorf("Expected 5 seats for a family car, instead got %d", familyCarSeats)
	}
	familyCarWheels := carVehicleF.NumWheels()
	if familyCarWheels != 4 {
		t.Errorf("Expected 4 wheels for a family car, instead got %d", familyCarWheels)
	}

	fCar, ok := carVehicleF.(Car)
	if !ok {
		t.Fatalf("Cast to car failed")
	}
	familyCarDoors := fCar.NumDoors()
	if familyCarDoors != 5 {
		t.Errorf("Expected 5 doors for a family car, instead got %d", familyCarDoors)
	}

	//LuxuryCar
	carVehicleL, lErr := carFactory.NewVehicle(LuxuryCarType)
	if lErr != nil {
		t.Fatal(lErr)
	}
	luxuryCarSeats := carVehicleL.NumSeats()
	if luxuryCarSeats != 2 {
		t.Errorf("Expected 2 seats for a luxury car, instead got %d", luxuryCarSeats)
	}
	luxuryCarWheels := carVehicleL.NumWheels()
	if luxuryCarWheels != 4 {
		t.Errorf("Expected 4 wheels for a luxury car, instead got %d", luxuryCarWheels)
	}

	lCar, ok := carVehicleL.(Car)
	if !ok {
		t.Errorf("Cast to car failed")
	}
	luxuryCarDoors := lCar.NumDoors()
	if luxuryCarDoors != 3 {
		t.Errorf("Expected 3 doors for a luxury car, instead got %d", luxuryCarDoors)
	}
}

func TestBuildFactoryMotorbike(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MotorbikeFactoryType)

	if err != nil {
		t.Fatal("Factory not recognized")
	}

	//SportMotorbike
	sportVehicleMotorbike, sErr := motorbikeFactory.NewVehicle(SportMotorbikeType)
	if sErr != nil {
		t.Fatal(sErr)
	}
	sMVSeats := sportVehicleMotorbike.NumSeats()
	if sMVSeats != 1 {
		t.Errorf("Expected 2 seats for sport motorbike, instead got %d", sMVSeats)
	}
	sMVWheels := sportVehicleMotorbike.NumWheels()
	if sMVWheels != 2 {
		t.Errorf("Expected 2 wheels for sport motorbike, instead got %d", sMVWheels)
	}

	sportMotorbike, ok := sportVehicleMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Cast to motorbike failed")
	}
	sMType := sportMotorbike.GetMotorbikeType()
	if sMType != 1 {
		t.Errorf("Expected sport type, instead got %d", sMType)
	}

	//CruiseMotorbike
	cruiseVehicleMotorbike, cErr := motorbikeFactory.NewVehicle(CruiseMotorbikeType)
	if cErr != nil {
		t.Fatal(cErr)
	}
	cMVSeats := cruiseVehicleMotorbike.NumSeats()
	if cMVSeats != 2 {
		t.Errorf("Expected 2 seats for cruise motorbike, instead got %d", cMVSeats)
	}
	cMVWheels := cruiseVehicleMotorbike.NumWheels()
	if cMVWheels != 2 {
		t.Errorf("Expected 2 wheels for cruise motorbike, instead got %d", cMVWheels)
	}

	cruiseMotorbike, ok := cruiseVehicleMotorbike.(Motorbike)
	if !ok {
		t.Fatal("Cast to motorbike failed")
	}
	cMType := cruiseMotorbike.GetMotorbikeType()
	if cMType != 2 {
		t.Errorf("Expected cruise type, instead got %d", cMType)
	}
}

func TestBuildFactoryNonExistent(t *testing.T) {
	_, err := BuildFactory(-1)
	if err == nil {
		t.Errorf("Expected Error, build factory does not exists")
	}
}

func TestBuildFactoryCarNonExistent(t *testing.T) {
	carFactory, _ := BuildFactory(CarFactoryType)
	_, cErr := carFactory.NewVehicle(-1)
	if cErr == nil {
		t.Errorf("Expected Error, build factory does not exists")
	}
}

func TestBuildFactoryMotorbikeNonExistent(t *testing.T) {
	motorbikeFactory, _ := BuildFactory(MotorbikeFactoryType)
	_, mErr := motorbikeFactory.NewVehicle(-1)
	if mErr == nil {
		t.Errorf("Expected Error, build factory does not exists")
	}
}