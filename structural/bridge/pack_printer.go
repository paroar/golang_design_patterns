package bridge

import "fmt"

type PackPrinter struct {
	Msg     string
	Printer IPrinterAPI
}

func (p *PackPrinter) Print() error {
	p.Printer.PrintMessage(fmt.Sprintf("Message from Pack: %s", p.Msg))
	return nil
}
