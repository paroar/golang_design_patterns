package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Fatal("A payment 'Cash' must exists")
	}

	msg := payment.Pay(10.5)
	if !strings.Contains(msg, "paid on cash") {
		t.Errorf("Cash payment message wasn't correct")
	}

	t.Log("LOG:", msg)
}
func TestGetPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Fatal("DebitCard payment message wasn't correct")
	}

	msg := payment.Pay(70.5)
	if !strings.Contains(msg, "paid on debit card") {
		t.Errorf("Expected DebitCardPM, instead got %s", "")
	}

	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(-1)

	if err == nil {
		t.Error("A payment method with ID 20 must return an error")
	}

	t.Log("LOG:", err)
}
