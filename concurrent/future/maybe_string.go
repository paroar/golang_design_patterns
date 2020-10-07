package future

import "fmt"

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (m *MaybeString) Success(f SuccessFunc) *MaybeString {
	m.successFunc = f
	return m
}
func (m *MaybeString) Fail(f FailFunc) *MaybeString {
	m.failFunc = f
	return m
}
func (m *MaybeString) Execute(f ExecuteStringFunc) {
	go func(m *MaybeString) {
		str, err := f()
		if err != nil {
			m.failFunc(err)
		} else {
			m.successFunc(str)
		}
	}(m)
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	return func() (string, error) {
		return msg, nil
	}
}
