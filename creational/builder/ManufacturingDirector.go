package builder


type ManufacturingDirector struct{
	builder Builder
}

func (f *ManufacturingDirector) SetBuilder(b Builder){
	f.builder = b
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetProduct().SetWheels().SetSeats()
}


