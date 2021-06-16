package interfaces

import (
	model "github.com/tavomartinez88/go-modules/orders/models"
)

type IOrder interface {
	Init()
	Build() model.Order
}