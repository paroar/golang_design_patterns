package main

import (
	"fmt"
	"os"
)

type AskState struct{}

func (a *AskState) executeState(c *Context) bool {

	fmt.Printf("Guess the number between 0 and 10, you have %d tries left\n", c.Retries)
	var guess int
	fmt.Fscanf(os.Stdin, "%d\n", &guess)

	if guess != c.SecretNumber {
		c.Retries -= 1
	} else {
		c.Won = true
		c.Next = &FinishState{}
	}

	if c.Retries < 1 {
		c.Won = false
		c.Next = &FinishState{}
	}

	return true
}
