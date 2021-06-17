package methodsConcretes

import "github.com/tavomartinez88/go-modules/orders/delivery/models"

type consumeOnSite struct {
	value string
}

const TypeDeliveryOnSite = "ON_SITE"

func (d *consumeOnSite) GetDeliveryMethod() models.Delivery {
	return models.Delivery{
		Type: TypeDeliveryOnSite,
		Description: d.value,
	}
}
