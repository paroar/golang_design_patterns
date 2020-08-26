package main

func main() {
	swimmer := CompositeSwimmer{  
		&TrainerImpl{},  
		&SwimmerImpl{},
	}
	swimmer.Swim()
	swimmer.Train()
}
