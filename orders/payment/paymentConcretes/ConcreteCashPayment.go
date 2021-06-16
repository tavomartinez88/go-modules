package paymentConcretes

import modelsPayment "github.com/tavomartinez88/go-modules/orders/payment/models"

type paymentCash struct {
	typePayment string
	value float64
}

func (c *paymentCash) SetType(wayPayment string) {
	c.typePayment = wayPayment
}

func (c *paymentCash) SetValue(value float64) {
	c.value = value
}

func (c *paymentCash) GetPayment() modelsPayment.Payment {
	return modelsPayment.Payment{
		Type: c.typePayment,
		PercentageForCard: nil,
		CountMoneyToPay: c.value,
	}
}

