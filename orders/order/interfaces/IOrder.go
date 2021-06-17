package interfaces

import (
	model "github.com/tavomartinez88/go-modules/orders/order/models"
)

type IOrder interface {
	Init()
	Build() (model.Order, error)
}