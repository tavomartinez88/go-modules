package paymentConcretes

import modelsPayment "github.com/tavomartinez88/go-modules/orders/payment/models"

type paymentCard struct {
	typePayment string
	value float64
}

func (c *paymentCard) GetPayment() modelsPayment.Payment {
	return modelsPayment.Payment{
		Type: c.typePayment,
		PercentageForCard: c.value,
	}
}

