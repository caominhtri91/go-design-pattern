package Factory

import (
	"fmt"
)

type CreditCardPM struct{}

func (d *CreditCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card (new)", amount)
}
