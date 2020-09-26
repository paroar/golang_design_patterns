package main

type OperationSum struct {
	Left  IInterpreter
	Right IInterpreter
}

func (o *OperationSum) Read() int {
	return o.Left.Read() + o.Right.Read()
}
