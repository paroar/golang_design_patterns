package chain_of_responsability

type ClosureChain struct {
	NextChain IChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}
	if c.NextChain != nil {
		c.NextChain.Next(s)
	}
}
