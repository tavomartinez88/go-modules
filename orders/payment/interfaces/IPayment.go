package interfaces

import (
	modelsPayment "github.com/tavomartinez88/go-modules/orders/payment/models"
)

type IPayment interface {
	GetPayment() modelsPayment.Payment
}
