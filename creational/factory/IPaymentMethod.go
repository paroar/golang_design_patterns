package factory

type IPaymentMethod interface {
	Pay(amount float64) string
}