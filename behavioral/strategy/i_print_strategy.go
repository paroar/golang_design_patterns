package strategy

import (
	"io"
)

type IPrintStrategy interface {
	Print() error
	SetLog(io.Writer)
	SetWriter(io.Writer)
}
