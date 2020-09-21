package chain_of_responsability

import (
	"fmt"
	"strings"
)

type SecondLogger struct {
	NextChain IChainLogger
}

func (sl *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if sl.NextChain != nil {
			sl.NextChain.Next(s)
		}

		return
	}

	fmt.Printf("Finishing in second logger\n")

}
