package chain_of_responsability

import "io"

type WriterLogger struct {
	NextChain IChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {

	if w.Writer != nil {
		w.Writer.Write([]byte("WriteLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}

}
