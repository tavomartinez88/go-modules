package interfaces

import (
	modelsPayment "github.com/tavomartinez88/go-modules/orders/payment/models"
)

type IPayment interface {
	SetType(wayPayment string)
	SetValue(value float64)
	GetPayment() modelsPayment.Payment
}
