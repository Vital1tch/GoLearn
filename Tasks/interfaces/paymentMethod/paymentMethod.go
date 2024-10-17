package paymentMethod

type PaymentMethod interface {
	Pay(amount float64) string
	CheckBalance() float64
}
