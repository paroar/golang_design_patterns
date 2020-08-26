package adapter

import "fmt"

type Adapter struct {
	OldPrinter *LegacyPrinter
	Msg        string
}

func (a *Adapter) PrintStored() (newMsg string) {
	if a.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", a.Msg)
		newMsg = a.OldPrinter.Print(newMsg)
	} else {
		newMsg = a.Msg
	}
	return
}
