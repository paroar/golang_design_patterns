package adapter

import "fmt"

type LegacyPrinter struct{}

func (*LegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	return
}
