package main

func main() {
	chatRoom := &ChatRoom{}

	user1 := &User{
		name:     "Pepe",
		mediator: chatRoom,
	}
	user2 := &User{
		name:     "Juan",
		mediator: chatRoom,
	}

	user1.SendMessage("Hello world")
	user2.SendMessage("Hi, hello")
}
