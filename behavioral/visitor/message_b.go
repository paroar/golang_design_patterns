package visitor

import (
	"fmt"
	"io"
	"os"
)

type MessageB struct {
	Msg    string
	Output io.Writer
}

func (m *MessageB) Accept(v IVisitor) {
	v.VisitB(m)
}

func (m *MessageB) Print() {
	if m.Output == nil {
		m.Output = os.Stdout
	}
	fmt.Fprintf(m.Output, "B: %s", m.Msg)
}
