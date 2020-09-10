package shapes

import "github.com/paroar/design_patterns/behavioral/strategy"

type TextSquare struct {
	strategy.PrintOutput
}

func (t *TextSquare) Print() error {
	t.Writer.Write([]byte("Square\n"))
	return nil
}
