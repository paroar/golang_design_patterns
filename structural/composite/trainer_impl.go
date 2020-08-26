package main

import "fmt"

type TrainerImpl struct{}
func (t *TrainerImpl) Train(){
	fmt.Println("Training")
}