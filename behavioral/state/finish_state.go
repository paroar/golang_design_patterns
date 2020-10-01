package main

type FinishState struct{}

func (f *FinishState) executeState(c *Context) bool {
	if c.Won {
		c.Next = &WinState{}
	} else {
		c.Next = &LoseState{}
	}
	return true
}
