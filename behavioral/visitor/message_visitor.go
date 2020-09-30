package visitor

import "fmt"

type MessageVisitor struct{}

func (mv *MessageVisitor) VisitA(m *MessageA) {
	m.Msg = fmt.Sprintf("%s%s", m.Msg, " (Visited A)")
}

func (mv *MessageVisitor) VisitB(m *MessageB) {
	m.Msg = fmt.Sprintf("%s%s", m.Msg, " (Visited B)")
}
