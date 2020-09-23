package template

import "strings"

type AnonTemplate struct{}

func (a *AnonTemplate) ExecuteAlgorithm(f func() string) string {
	strs := []string{a.first(), f(), a.third()}
	return strings.Join(strs, " ")
}

func (a *AnonTemplate) first() string {
	return "First"
}

func (a *AnonTemplate) third() string {
	return "Third"
}
