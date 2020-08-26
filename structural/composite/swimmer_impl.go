package main

import "fmt"

type SwimmerImpl struct{}
func (s *SwimmerImpl) Swim(){
	fmt.Println("Swimming")
}