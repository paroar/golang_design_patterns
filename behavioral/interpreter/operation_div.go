package main

type OperationDiv struct {
	Left  IInterpreter
	Right IInterpreter
}

func (o *OperationDiv) Read() int {
	return o.Left.Read() / o.Right.Read()
}
