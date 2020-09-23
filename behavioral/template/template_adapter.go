package template

type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	return ""
}

func MessageRetrieverAdapter(f func() string) IMessageRetriever {
	return &adapter{myFunc: f}
}
