package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type StartState struct{}

func (s *StartState) executeState(c *Context) bool {
	c.Next = &AskState{}

	rand.Seed(time.Now().UnixNano())
	c.SecretNumber = rand.Intn(10)

	fmt.Println("Introduce a number of retries:")
	fmt.Fscanf(os.Stdin, "%d\n", &c.Retries)

	return true
}
