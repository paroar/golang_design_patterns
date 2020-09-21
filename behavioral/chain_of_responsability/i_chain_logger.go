package chain_of_responsability

type IChainLogger interface {
	Next(string)
}
