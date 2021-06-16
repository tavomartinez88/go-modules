package paymentConcretes

import "github.com/tavomartinez88/go-modules/orders/payment/interfaces"

const (
	CASH = "CASH"
)

type cash struct {
	paymentCash
}

func CreateCashCreator(v float64) interfaces.IPayment {
	return &cash{
		paymentCash: paymentCash{
			typePayment: CASH,
			value: v,
		},
	}
}
