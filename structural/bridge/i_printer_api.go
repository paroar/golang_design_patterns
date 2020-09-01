package bridge

type IPrinterAPI interface {
	PrintMessage(string) error
}
