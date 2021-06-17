package paymentConcretes

import modelsPayment "github.com/tavomartinez88/go-modules/orders/payment/models"

type paymentCash struct {
	typePayment string
	value float64
}

func (c *paymentCash) GetPayment() modelsPayment.Payment {
	return modelsPayment.Payment{
		Type: c.typePayment,
		CountMoneyToPay: c.value,
	}
}

