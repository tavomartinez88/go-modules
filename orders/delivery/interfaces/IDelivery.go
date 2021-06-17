package interfaces

import "github.com/tavomartinez88/go-modules/orders/delivery/models"

type IDelivery interface {
	GetDeliveryMethod() models.Delivery
}