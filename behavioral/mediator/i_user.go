package main

type IUser interface {
	SendMessage(string)
	GetName() string
}
