package strategy

import (
	"io"
)

type PrintOutput struct {
	Writer    io.Writer
	LogWriter io.Writer
}

func (p *PrintOutput) SetLog(w io.Writer) {
	p.LogWriter = w
}
func (p *PrintOutput) SetWriter(w io.Writer) {
	p.Writer = w
}
