package prototype

import "fmt"

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Price: %.2f\nSKU: %s\nColor id: %d\n", s.Price, s.SKU, s.Color)
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
