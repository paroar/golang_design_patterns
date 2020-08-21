package factory

import "fmt"

type CashPM struct{}

func (c *CashPM) Pay(amount float64) string{
	return fmt.Sprintf("paid on cash %0.2f", amount)
}