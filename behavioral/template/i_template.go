package template

type ITemplate interface {
	first()
	third()
	ExecuteAlgorithm(IMessageRetriever) string
}
