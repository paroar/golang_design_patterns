package main

import "fmt"

func main() {
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(true))
	assertAndPrint(m.RestoreSettings())
	assertAndPrint(m.RestoreSettings())
}

func assertAndPrint(c ICommand) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
