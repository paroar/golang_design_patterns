package main

type MementoFacade struct {
	originator originator
	careTaker  careTaker
}

func (m *MementoFacade) SaveSettings(s ICommand) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings() ICommand {
	m.originator.ExtracAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}
