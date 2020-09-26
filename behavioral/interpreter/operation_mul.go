package main

type OperationMul struct {
	Left  IInterpreter
	Right IInterpreter
}

func (o *OperationMul) Read() int {
	return o.Left.Read() * o.Right.Read()
}
