package methodsConcretes

import "github.com/tavomartinez88/go-modules/orders/delivery/models"

type delivery struct {
	value string
}

const TypeDelivery = "DELIVERY"

func (d *delivery) GetDeliveryMethod() models.Delivery {
	return models.Delivery{
		Type: TypeDelivery,
		Description: d.value,
	}
}
