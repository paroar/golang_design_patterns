package main

import (
	"fmt"
	"strings"
)

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func (p *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	p.prefix(p.append(p.uppercase(in)))
}

func (p *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (p *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", s, p.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (p *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercasedStringWithSuffix, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", p.prefixS, uppercasedStringWithSuffix))
		}
	}()
}
