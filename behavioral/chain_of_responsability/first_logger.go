package chain_of_responsability

import "fmt"

type FirstLogger struct {
	NextChain IChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}
