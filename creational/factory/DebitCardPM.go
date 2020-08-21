package factory

import "fmt"

type DebitCardPM struct{}

func (d *DebitCardPM) Pay(amount float64) string{
	return fmt.Sprintf("paid on debit card %0.2f", amount)
}