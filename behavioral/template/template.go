package template

import "strings"

type Template struct{}

func (t *Template) ExecuteAlgorithm(m IMessageRetriever) string {
	strs := []string{t.first(), m.Message(), t.third()}
	return strings.Join(strs, " ")
}

func (t *Template) first() string {
	return "First"
}

func (t *Template) third() string {
	return "Third"
}
