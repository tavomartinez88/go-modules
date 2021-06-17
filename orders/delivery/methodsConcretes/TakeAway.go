package methodsConcretes

import "github.com/tavomartinez88/go-modules/orders/delivery/models"

type takeAway struct {
}

const TypeTakeAWay = "TAKE_A_WAY"
const description = "lo retiro personalmente"

func (d *takeAway) GetDeliveryMethod() models.Delivery {
	return models.Delivery{
		Type:        TypeTakeAWay,
		Description: description,
	}
}
