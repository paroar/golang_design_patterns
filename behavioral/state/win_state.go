package main

type WinState struct{}

func (w *WinState) executeState(c *Context) bool {
	println("Congrats, you won")
	return false
}
