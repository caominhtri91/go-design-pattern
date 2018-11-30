package Factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't corrent")
	}

	t.Log("Log:", msg)
}

func TestCreatePaymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(Debit)
	if err != nil {
		t.Fatal("A payment method of type 'Debit' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit") {
		t.Error("The debit payment method message wasn't correct")
	}

	t.Log("Log:", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(20)
	if err != nil {
		t.Error("A payment method with ID 20 must return an error")
	}

	t.Log("Log:", err)
}
