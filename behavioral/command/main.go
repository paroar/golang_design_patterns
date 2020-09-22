package main

import (
	"fmt"
	"time"
)

func main() {
	third := Logger{}
	second := Logger{NextChain: &third}
	first := Logger{NextChain: &second}

	command := &TimePassed{start: time.Now()}

	first.Next(command)
}

type ICommand interface {
	Info() string
}

type TimePassed struct {
	start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.start).String()
}

type IChainLogger interface {
	Next(ICommand)
}

type Logger struct {
	NextChain IChainLogger
}

func (l *Logger) Next(c ICommand) {
	time.Sleep(time.Second)

	fmt.Printf("Elapsed time from creation: %s\n", c.Info())
	if l.NextChain != nil {
		l.NextChain.Next(c)
	}
}
