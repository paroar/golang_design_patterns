package main

import (
	"strconv"
	"strings"
)

func main() {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub 8 sub", " ")
	for _, operatorString := range operators {
		if IsOperator(operatorString) {
			right := stack.Pop()
			left := stack.Pop()
			mathFunc := operatorFactory(operatorString, left, right)
			res := value(mathFunc.Read())
			stack.Push(&res)
		} else {
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}
	println(int(stack.Pop().Read()))
}
