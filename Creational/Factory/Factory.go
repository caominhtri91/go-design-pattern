package Factory

import (
	"fmt"
)

const (
	Cash  = 1
	Debit = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case Debit:
		return new(CreditCardPM), nil
	default:
		return nil, fmt.Errorf("payment method %d not recognized", m)
	}
}
