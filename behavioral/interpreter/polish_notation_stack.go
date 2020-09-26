package main

type polishNotationStack []IInterpreter

func (p *polishNotationStack) Push(s IInterpreter) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() IInterpreter {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return nil
}
