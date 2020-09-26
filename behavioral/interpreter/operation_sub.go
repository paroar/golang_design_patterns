package main

type OperationSub struct {
	Left  IInterpreter
	Right IInterpreter
}

func (o *OperationSub) Read() int {
	return o.Left.Read() - o.Right.Read()
}
