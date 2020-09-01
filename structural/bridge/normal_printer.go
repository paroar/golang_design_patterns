package bridge

type NormalPrinter struct {
	Msg     string
	Printer IPrinterAPI
}

func (p *NormalPrinter) Print() error {
	p.Printer.PrintMessage(p.Msg)
	return nil
}
