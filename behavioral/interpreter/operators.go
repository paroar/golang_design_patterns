package main

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func operatorFactory(o string, left, right IInterpreter) IInterpreter {
	switch o {
	case SUM:
		return &OperationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &OperationSub{
			Left:  left,
			Right: right,
		}

	case MUL:
		return &OperationMul{
			Left:  left,
			Right: right,
		}
	case DIV:
		return &OperationDiv{
			Left:  left,
			Right: right,
		}
	}
	return nil
}

func IsOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}
	return false
}
