package main

type User struct {
	name     string
	mediator IMediator
}

func (u *User) SendMessage(msg string) {
	u.mediator.ShowMessage(u, msg)
}

func (u *User) GetName() string {
	return u.name
}
