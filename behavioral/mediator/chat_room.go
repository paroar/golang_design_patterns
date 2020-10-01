package main

import "fmt"

type ChatRoom struct{}

func (c *ChatRoom) ShowMessage(u IUser, msg string) {
	fmt.Printf("%s: %s\n", u.GetName(), msg)
}
