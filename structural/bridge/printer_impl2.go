package bridge

import (
	"errors"
	"fmt"
	"io"
)

type PrinterImpl2 struct {
	Writer io.Writer
}

func (p *PrinterImpl2) PrintMessage(msg string) error {
	if p.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(p.Writer, "%s", msg)
	return nil
}
