package main

type originator struct {
	Command ICommand
}

func (o *originator) NewMemento() memento {
	return memento{Command: o.Command}
}

func (o *originator) ExtracAndStoreCommand(m memento) {
	o.Command = m.Command
}
