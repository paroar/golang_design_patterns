package main

import "fmt"

type LoseState struct{}

func (w *LoseState) executeState(c *Context) bool {
	fmt.Printf("You lose. The correct number was: %d\n", c.SecretNumber)
	return false
}
