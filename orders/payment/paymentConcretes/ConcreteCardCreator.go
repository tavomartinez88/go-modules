package paymentConcretes

import "github.com/tavomartinez88/go-modules/orders/payment/interfaces"

const (
	CARD = "CARD"
)

type card struct {
	paymentCard
}

func CreateCardPayment(v float64) interfaces.IPayment {
	return &card{
		paymentCard: paymentCard{
			typePayment: CARD,
			value: v,
		},
	}
}
