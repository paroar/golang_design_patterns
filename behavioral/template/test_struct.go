package template

type TestStruct struct{}

func (t *TestStruct) Message() string {
	return "world"
}
