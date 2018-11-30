package Factory

type PaymentMethod interface {
	Pay(amount float32) string
}
