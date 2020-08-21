package factory

import (
	"errors"
	"fmt"
)

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (IPaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized", m))
	}
}
